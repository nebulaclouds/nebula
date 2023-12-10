package tasklog

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
)

func MustCreateRegex(varName string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`(?i){{\s*[\.$]%s\s*}}`, varName))
}

type templateRegexes struct {
	LogName              *regexp.Regexp
	PodName              *regexp.Regexp
	PodUID               *regexp.Regexp
	Namespace            *regexp.Regexp
	ContainerName        *regexp.Regexp
	ContainerID          *regexp.Regexp
	Hostname             *regexp.Regexp
	PodRFC3339StartTime  *regexp.Regexp
	PodRFC3339FinishTime *regexp.Regexp
	PodUnixStartTime     *regexp.Regexp
	PodUnixFinishTime    *regexp.Regexp
	TaskID               *regexp.Regexp
	TaskVersion          *regexp.Regexp
	TaskProject          *regexp.Regexp
	TaskDomain           *regexp.Regexp
	TaskRetryAttempt     *regexp.Regexp
	NodeID               *regexp.Regexp
	ExecutionName        *regexp.Regexp
	ExecutionProject     *regexp.Regexp
	ExecutionDomain      *regexp.Regexp
	GeneratedName        *regexp.Regexp
}

func initDefaultRegexes() templateRegexes {
	return templateRegexes{
		MustCreateRegex("logName"),
		MustCreateRegex("podName"),
		MustCreateRegex("podUID"),
		MustCreateRegex("namespace"),
		MustCreateRegex("containerName"),
		MustCreateRegex("containerID"),
		MustCreateRegex("hostname"),
		MustCreateRegex("podRFC3339StartTime"),
		MustCreateRegex("podRFC3339FinishTime"),
		MustCreateRegex("podUnixStartTime"),
		MustCreateRegex("podUnixFinishTime"),
		MustCreateRegex("taskID"),
		MustCreateRegex("taskVersion"),
		MustCreateRegex("taskProject"),
		MustCreateRegex("taskDomain"),
		MustCreateRegex("taskRetryAttempt"),
		MustCreateRegex("nodeID"),
		MustCreateRegex("executionName"),
		MustCreateRegex("executionProject"),
		MustCreateRegex("executionDomain"),
		MustCreateRegex("generatedName"),
	}
}

var defaultRegexes = initDefaultRegexes()

func replaceAll(template string, vars TemplateVars) string {
	for _, v := range vars {
		if len(v.Value) > 0 {
			template = v.Regex.ReplaceAllLiteralString(template, v.Value)
		}
	}
	return template
}

func (input Input) templateVarsForScheme(scheme TemplateScheme) TemplateVars {
	vars := TemplateVars{
		{defaultRegexes.LogName, input.LogName},
	}

	gotExtraTemplateVars := input.ExtraTemplateVarsByScheme != nil
	if gotExtraTemplateVars {
		vars = append(vars, input.ExtraTemplateVarsByScheme.Common...)
	}

	switch scheme {
	case TemplateSchemePod:
		// Container IDs are prefixed with docker://, cri-o://, etc. which is stripped by fluentd before pushing to a log
		// stream. Therefore, we must also strip the prefix.
		containerID := input.ContainerID
		stripDelimiter := "://"
		if split := strings.Split(input.ContainerID, stripDelimiter); len(split) > 1 {
			containerID = split[1]
		}
		vars = append(
			vars,
			TemplateVar{defaultRegexes.PodName, input.PodName},
			TemplateVar{defaultRegexes.PodUID, input.PodUID},
			TemplateVar{defaultRegexes.Namespace, input.Namespace},
			TemplateVar{defaultRegexes.ContainerName, input.ContainerName},
			TemplateVar{defaultRegexes.ContainerID, containerID},
			TemplateVar{defaultRegexes.Hostname, input.HostName},
			TemplateVar{defaultRegexes.PodRFC3339StartTime, input.PodRFC3339StartTime},
			TemplateVar{defaultRegexes.PodRFC3339FinishTime, input.PodRFC3339FinishTime},
			TemplateVar{
				defaultRegexes.PodUnixStartTime,
				strconv.FormatInt(input.PodUnixStartTime, 10),
			},
			TemplateVar{
				defaultRegexes.PodUnixFinishTime,
				strconv.FormatInt(input.PodUnixFinishTime, 10),
			},
		)
		if gotExtraTemplateVars {
			vars = append(vars, input.ExtraTemplateVarsByScheme.Pod...)
		}
	case TemplateSchemeTaskExecution:
		taskExecutionIdentifier := input.TaskExecutionID.GetID()
		vars = append(
			vars,
			TemplateVar{
				defaultRegexes.NodeID,
				input.TaskExecutionID.GetUniqueNodeID(),
			},
			TemplateVar{
				defaultRegexes.GeneratedName,
				input.TaskExecutionID.GetGeneratedName(),
			},
			TemplateVar{
				defaultRegexes.TaskRetryAttempt,
				strconv.FormatUint(uint64(taskExecutionIdentifier.RetryAttempt), 10),
			},
		)
		if taskExecutionIdentifier.TaskId != nil {
			vars = append(
				vars,
				TemplateVar{
					defaultRegexes.TaskID,
					taskExecutionIdentifier.TaskId.Name,
				},
				TemplateVar{
					defaultRegexes.TaskVersion,
					taskExecutionIdentifier.TaskId.Version,
				},
				TemplateVar{
					defaultRegexes.TaskProject,
					taskExecutionIdentifier.TaskId.Project,
				},
				TemplateVar{
					defaultRegexes.TaskDomain,
					taskExecutionIdentifier.TaskId.Domain,
				},
			)
		}
		if taskExecutionIdentifier.NodeExecutionId != nil && taskExecutionIdentifier.NodeExecutionId.ExecutionId != nil {
			vars = append(
				vars,
				TemplateVar{
					defaultRegexes.ExecutionName,
					taskExecutionIdentifier.NodeExecutionId.ExecutionId.Name,
				},
				TemplateVar{
					defaultRegexes.ExecutionProject,
					taskExecutionIdentifier.NodeExecutionId.ExecutionId.Project,
				},
				TemplateVar{
					defaultRegexes.ExecutionDomain,
					taskExecutionIdentifier.NodeExecutionId.ExecutionId.Domain,
				},
			)
		}
		if gotExtraTemplateVars {
			vars = append(vars, input.ExtraTemplateVarsByScheme.TaskExecution...)
		}
	}

	return vars
}

func (p TemplateLogPlugin) GetTaskLogs(input Input) (Output, error) {
	templateVars := input.templateVarsForScheme(p.Scheme)
	taskLogs := make([]*core.TaskLog, 0, len(p.TemplateURIs))
	for _, templateURI := range p.TemplateURIs {
		taskLogs = append(taskLogs, &core.TaskLog{
			Uri:           replaceAll(templateURI, templateVars),
			Name:          p.DisplayName + input.LogName,
			MessageFormat: p.MessageFormat,
		})
	}

	return Output{TaskLogs: taskLogs}, nil
}
