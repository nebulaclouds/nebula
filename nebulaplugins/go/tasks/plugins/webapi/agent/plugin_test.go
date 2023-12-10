package agent

import (
	"context"
	nebulaidlcore "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	pluginsCore "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core"
	pluginCoreMocks "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core/mocks"
	webapiPlugin "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/webapi/mocks"
	"github.com/nebulaclouds/nebula/nebulastdlib/config"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

func TestPlugin(t *testing.T) {
	fakeSetupContext := pluginCoreMocks.SetupContext{}
	fakeSetupContext.OnMetricsScope().Return(promutils.NewScope("test"))

	cfg := defaultConfig
	cfg.WebAPI.Caching.Workers = 1
	cfg.WebAPI.Caching.ResyncInterval.Duration = 5 * time.Second
	cfg.DefaultAgent = Agent{Endpoint: "test-agent.nebula.svc.cluster.local:80"}
	cfg.Agents = map[string]*Agent{"spark_agent": {Endpoint: "localhost:80"}}
	cfg.AgentForTaskTypes = map[string]string{"spark": "spark_agent", "bar": "bar_agent"}

	plugin := Plugin{
		metricScope: fakeSetupContext.MetricsScope(),
		cfg:         GetConfig(),
	}
	t.Run("get config", func(t *testing.T) {
		err := SetConfig(&cfg)
		assert.NoError(t, err)
		assert.Equal(t, cfg.WebAPI, plugin.GetConfig())
	})
	t.Run("get ResourceRequirements", func(t *testing.T) {
		namespace, constraints, err := plugin.ResourceRequirements(context.TODO(), nil)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.ResourceNamespace("default"), namespace)
		assert.Equal(t, plugin.cfg.ResourceConstraints, constraints)
	})

	t.Run("test newAgentPlugin", func(t *testing.T) {
		p := newAgentPlugin()
		assert.NotNil(t, p)
		assert.Equal(t, "agent-service", p.ID)
		assert.NotNil(t, p.PluginLoader)
	})

	t.Run("test getFinalAgent", func(t *testing.T) {
		agent, _ := getFinalAgent("spark", &cfg)
		assert.Equal(t, cfg.Agents["spark_agent"].Endpoint, agent.Endpoint)
		agent, _ = getFinalAgent("foo", &cfg)
		assert.Equal(t, cfg.DefaultAgent.Endpoint, agent.Endpoint)
		_, err := getFinalAgent("bar", &cfg)
		assert.NotNil(t, err)
	})

	t.Run("test getClientFunc", func(t *testing.T) {
		client, err := getClientFunc(context.Background(), &Agent{Endpoint: "localhost:80"}, map[*Agent]*grpc.ClientConn{})
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("test getClientFunc more config", func(t *testing.T) {
		client, err := getClientFunc(context.Background(), &Agent{Endpoint: "localhost:80", Insecure: true, DefaultServiceConfig: "{\"loadBalancingConfig\": [{\"round_robin\":{}}]}"}, map[*Agent]*grpc.ClientConn{})
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("test getClientFunc cache hit", func(t *testing.T) {
		connectionCache := make(map[*Agent]*grpc.ClientConn)
		agent := &Agent{Endpoint: "localhost:80", Insecure: true, DefaultServiceConfig: "{\"loadBalancingConfig\": [{\"round_robin\":{}}]}"}

		client, err := getClientFunc(context.Background(), agent, connectionCache)
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.NotNil(t, client, connectionCache[agent])

		cachedClient, err := getClientFunc(context.Background(), agent, connectionCache)
		assert.NoError(t, err)
		assert.NotNil(t, cachedClient)
		assert.Equal(t, client, cachedClient)
	})

	t.Run("test getFinalTimeout", func(t *testing.T) {
		timeout := getFinalTimeout("CreateTask", &Agent{Endpoint: "localhost:8080", Timeouts: map[string]config.Duration{"CreateTask": {Duration: 1 * time.Millisecond}}})
		assert.Equal(t, 1*time.Millisecond, timeout.Duration)
		timeout = getFinalTimeout("DeleteTask", &Agent{Endpoint: "localhost:8080", DefaultTimeout: config.Duration{Duration: 10 * time.Second}})
		assert.Equal(t, 10*time.Second, timeout.Duration)
	})

	t.Run("test getFinalContext", func(t *testing.T) {
		ctx, _ := getFinalContext(context.TODO(), "DeleteTask", &Agent{})
		assert.Equal(t, context.TODO(), ctx)

		ctx, _ = getFinalContext(context.TODO(), "CreateTask", &Agent{Endpoint: "localhost:8080", Timeouts: map[string]config.Duration{"CreateTask": {Duration: 1 * time.Millisecond}}})
		assert.NotEqual(t, context.TODO(), ctx)
	})

	t.Run("test PENDING Status", func(t *testing.T) {
		taskContext := new(webapiPlugin.StatusContext)
		taskContext.On("Resource").Return(ResourceWrapper{
			State:    admin.State_PENDING,
			Outputs:  nil,
			Message:  "Waiting for cluster",
			LogLinks: []*nebulaidlcore.TaskLog{{Uri: "http://localhost:3000/log", Name: "Log Link"}},
		})

		phase, err := plugin.Status(context.Background(), taskContext)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseInitializing, phase.Phase())
		assert.Equal(t, "Waiting for cluster", phase.Reason())
	})

	t.Run("test RUNNING Status", func(t *testing.T) {
		taskContext := new(webapiPlugin.StatusContext)
		taskContext.On("Resource").Return(ResourceWrapper{
			State:    admin.State_RUNNING,
			Outputs:  nil,
			Message:  "Job is running",
			LogLinks: []*nebulaidlcore.TaskLog{{Uri: "http://localhost:3000/log", Name: "Log Link"}},
		})

		phase, err := plugin.Status(context.Background(), taskContext)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseRunning, phase.Phase())
	})

	t.Run("test PERMANENT_FAILURE Status", func(t *testing.T) {
		taskContext := new(webapiPlugin.StatusContext)
		taskContext.On("Resource").Return(ResourceWrapper{
			State:    admin.State_PERMANENT_FAILURE,
			Outputs:  nil,
			Message:  "",
			LogLinks: []*nebulaidlcore.TaskLog{{Uri: "http://localhost:3000/log", Name: "Log Link"}},
		})

		phase, err := plugin.Status(context.Background(), taskContext)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhasePermanentFailure, phase.Phase())
	})

	t.Run("test RETRYABLE_FAILURE Status", func(t *testing.T) {
		taskContext := new(webapiPlugin.StatusContext)
		taskContext.On("Resource").Return(ResourceWrapper{
			State:    admin.State_RETRYABLE_FAILURE,
			Outputs:  nil,
			Message:  "",
			LogLinks: []*nebulaidlcore.TaskLog{{Uri: "http://localhost:3000/log", Name: "Log Link"}},
		})

		phase, err := plugin.Status(context.Background(), taskContext)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseRetryableFailure, phase.Phase())
	})

	t.Run("test UNDEFINED Status", func(t *testing.T) {
		taskContext := new(webapiPlugin.StatusContext)
		taskContext.On("Resource").Return(ResourceWrapper{
			State:    5,
			Outputs:  nil,
			Message:  "",
			LogLinks: []*nebulaidlcore.TaskLog{{Uri: "http://localhost:3000/log", Name: "Log Link"}},
		})

		phase, err := plugin.Status(context.Background(), taskContext)
		assert.Error(t, err)
		assert.Equal(t, pluginsCore.PhaseUndefined, phase.Phase())
	})
}
