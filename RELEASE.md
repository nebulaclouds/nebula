# Release Process

[![hackmd-github-sync-badge](https://hackmd.io/sVOAyv6LTwiQllQUctxP1w/badge)](https://hackmd.io/sVOAyv6LTwiQllQUctxP1w)

## Resolve Issues
1. Create [a new milestone](https://github.com/nebulaclouds/nebula/milestones) if one doesn't exist.
1. Open [issues](https://github.com/nebulaclouds/nebula/issues) and filter by milestone and make sure they are either closed or moved over to the next milestone.
## Start a release PR
1. Run [Generate Nebula Manifests workflow](https://github.com/nebulaclouds/nebula/actions/workflows/generate_nebula_manifest.yml). It’ll create a PR ([example](https://github.com/nebulaclouds/nebula/pull/888))
1. Update [docs version](https://github.com/nebulaclouds/nebula/blob/master/docs/conf.py#L33) to match the milestone version.
1. Create a CHANGELOG file ([example](https://github.com/nebulaclouds/nebula/pull/888/files#diff-0c33dda4ecbd7e1116ddce683b5e143d85b22e43223ca258ecc571fb3b240a57))
1. Wait for endtoend tests to finish then Merge PR.
## Create a release
1. Run [Create Nebula Release workflow](https://github.com/nebulaclouds/nebula/actions/workflows/create_release.yml):
   It will create a tag and then publish all deployment manifest in github release and will create a discussion thread in github release
1. Kick off a run of the functional tests in https://github.com/unionai/genesis-device/actions/workflows/update_cluster_and_run_tests.yml
1. Close the milestone
1. Ping #core ([slack](https://slack.nebula.org/) channel) to send announcements about the milestone with the contents of the CHANGELOG to all social channels.
