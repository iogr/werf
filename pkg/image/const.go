package image

const (
	WerfLabel                      = "werf"
	WerfVersionLabel               = "werf-version"
	WerfCacheVersionLabel          = "werf-cache-version"
	WerfImageLabel                 = "werf-image"
	WerfImageNameLabel             = "werf-image-name"
	WerfDockerImageName            = "werf-docker-image-name"
	WerfStageSignatureLabel        = "werf-stage-signature"
	WerfStageContentSignatureLabel = "werf-stage-content-signature"
	WerfProjectRepoCommitLabel     = "werf-project-repo-commit"
	WerfContentSignatureLabel      = "werf-content-signature"

	WerfMountTmpDirLabel          = "werf-mount-type-tmp-dir"
	WerfMountBuildDirLabel        = "werf-mount-type-build-dir"
	WerfMountCustomDirLabelPrefix = "werf-mount-type-custom-dir-"

	WerfImportLabelPrefix = "werf-import-"

	BuildCacheVersion = "1.1"

	StageContainerNamePrefix = "werf.build."
)
