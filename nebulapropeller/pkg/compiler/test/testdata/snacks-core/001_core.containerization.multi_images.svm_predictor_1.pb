
�
4"0core.containerization.multi_images.svm_predictorpython-task0.32.6python* "�
�

y_test
B	parquety_test
!
y_train
B	parquety_train
!
X_train
B	parquetX_train

X_test
B	parquetX_test

o0
o02�
sghcr.io/nebulaclouds/nebulacookbook:multi-image-predict-98b125fd57d20594026941c2ebe7ef662e5acb7d6423660a65f493ca2d9aa267pynebula-execute--inputs
{{.input}}--output-prefix{{.outputPrefix}}--raw-output-data-prefix{{.rawOutputDataPrefix}}--checkpoint-path{{.checkpointOutputPrefix}}--prev-checkpoint{{.prevCheckpointPrefix}}
--resolver9nebulakit.core.python_auto_container.default_task_resolver--task-module"core.containerization.multi_images	task-name
svm_predictor" 