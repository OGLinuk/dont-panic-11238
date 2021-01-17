package dontpanic

var (
	// DefaultEnvs is a map of default environments and their default manifests definitions
	DefaultEnvs = map[string]map[string][]string{
		"default": map[string][]string{
			"functions": []string{
				"sbh 9001 https://gitlab.com/oglinuk/sbh.git",
				"quotitioner 9429 https://gitlab.com/oglinuk/quotitioner.git",
			},
			"applications": []string{
				// TODO: Fix (to 9999) when gendockercompose is implemented
				"gitea 3000 https://github.com/go-gitea/gitea.git",
			},
			"individuals-blog": []string{
				"fragglet-blog 50006 https://github.com/fragglet/soulsphere.org.git",
				"rwxrob-blog 50007 https://github.com/rwxrob/rwxrob.git",
				"jessfraz-blog 50008 https://github.com/jessfraz/blog.git",
			},
			"projects-blog": []string{
				// TODO: Uncomment when the ability to dynamically change port is implemented
				//"golang-blog 50300 https://github.com/golang/blog.git",
				"gitea-blog 50301 https://gitea.com/gitea/blog.git",
			},
			"docs": []string{
				"git-docs 50500 https://github.com/git/htmldocs.git",
				"rwx.gg 50501 https://gitlab.com/rwx.gg/README.git",
				"freedoom-docs 50502 https://github.com/freedoom/freedoom.github.io.git",
			},
			"fileserver": []string{
				"mediafs 12121 https://gitlab.com/oglinuk/mediafs.git",
				"directories 50000 https://gitlab.com/oglinuk/directories.git",
				"library 50001 https://gitlab.com/oglinuk/library.git",
				"ptp 50002 https://github.com/oglinuk/ptp.git",
				// TODO: Uncomment when `pandoc.go` is implemented
				// "awesome-gitea 50700 https://gitea.com/gitea/awesome-gitea.git",
			},
			"games": []string{
				"chocolate-doom 30303 https://github.com/chocolate-doom/chocolate-doom.git",
				"quakespasm 30304 https://github.com/ericwa/quakespasm.git",
			},
			"personal": []string{
				//"fourohfournotfound 40404 https://gitlab.com/oglinuk/fourohfournotfound.git"
			},
		},
	}
)
