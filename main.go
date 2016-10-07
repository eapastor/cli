package main

import (
	"os"
	"github.com/lastbackend/cli/cmd"
)



//"BITBUCKET ssh":   "ssh://hg@bitbucket.example.com:7999/PROJ/repo",

func main() {

	//giturls := []string{
	//	"https://github.com/lastbackend/vendors.git",
	//	"git@github.com:lastbackend/vendors.git",
	//	"git@bitbucket.org:p_evgenii/c.git",
	//	"https://p_evgenii@bitbucket.org/p_evgenii/c.git",
	//	"https://gitlab.com/e.pastor2926/test.git",
	//	"git@gitlab.com:e.pastor2926/test.git",
	//	//"git@github.com:user/project.git",
	//	//"https://github.com/user/project.git",
	//	//"http://github.com/user/project.git",
	//	//"git@192.168.101.127:user/project.git",
	//	//"https://192.168.101.127/user/project.git",
	//	//"http://192.168.101.127/user/project.git",
	//	//"ssh://user@host.xz:port/path/to/repo.git/",
	//	//"ssh://user@host.xz/path/to/repo.git/",
	//	//"ssh://host.xz:port/path/to/repo.git/",
	//	//"ssh://host.xz/path/to/repo.git/",
	//	//"ssh://user@host.xz/path/to/repo.git/",
	//	//"ssh://host.xz/path/to/repo.git/",
	//	//"ssh://user@host.xz/~user/path/to/repo.git/",
	//	//"ssh://host.xz/~user/path/to/repo.git/",
	//	//"ssh://user@host.xz/~/path/to/repo.git",
	//	//"ssh://host.xz/~/path/to/repo.git",
	//	//"git://host.xz/path/to/repo.git/",
	//	//"git://host.xz/~user/path/to/repo.git/",
	//	//"http://host.xz/path/to/repo.git/",
	//	//"https://host.xz/path/to/repo.git/",
	//	//"/path/to/repo.git/",
	//	//"path/to/repo.git/",
	//	//"~/path/to/repo.git",
	//	//"file:///path/to/repo.git/",
	//	//"file://~/path/to/repo.git/",
	//	//"user@host.xz:/path/to/repo.git/",
	//	//"host.xz:/path/to/repo.git/",
	//	//"user@host.xz:~user/path/to/repo.git/",
	//	//"host.xz:~user/path/to/repo.git/",
	//	//"user@host.xz:path/to/repo.git",
	//	//"host.xz:path/to/repo.git",
	//	//"rsync://host.xz/path/to/repo.git/",
	//}
	//
	//for _, v := range giturls {
	//	fmt.Printf("\n \x1b[32m %s: \x1b[35m \n", v)
	//	fmt.Printf("\n \x1b[33m %+v: \x1b[39m \n", utils.GitUrlParse(v))
	//}

	cmd.Execute()
	os.Exit(0)

}
