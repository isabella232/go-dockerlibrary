package manifest_test

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/docker-library/go-dockerlibrary/manifest"
)

func Example() {
	man, err := manifest.Parse(bufio.NewReader(strings.NewReader(`# RFC 2822

	# I LOVE CAKE

Maintainers: InfoSiftr <github@infosiftr.com> (@infosiftr),
             Johan Euphrosine <proppy@google.com> (@proppy)
GitRepo: https://github.com/docker-library/golang.git
GitFetch: refs/heads/master
SharedTags: latest


 # hi


 	 # blasphemer


# Go 1.6
Tags: 1.6.1, 1.6, 1
GitCommit: 0ce80411b9f41e9c3a21fc0a1bffba6ae761825a
Directory: 1.6


# Go 1.5
Tags: 1.5.3
SharedTags: 1.5.3-debian, 1.5-debian
GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
Directory: 1.5


Tags: 1.5
SharedTags: 1.5-debian
GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
Directory: 1.5


`)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("-------------\n2822:\n%s\n", man)

	fmt.Printf("\nShared Tag Groups:\n")
	for _, group := range man.GetSharedTagGroups() {
		fmt.Printf("\n  - %s\n", strings.Join(group.SharedTags, ", "))
		for _, entry := range group.Entries {
			fmt.Printf("    - %s\n", entry.TagsString())
		}
	}
	fmt.Printf("\n")

	man, err = manifest.Parse(bufio.NewReader(strings.NewReader(`
# maintainer: InfoSiftr <github@infosiftr.com> (@infosiftr)
# maintainer: John Smith <jsmith@example.com> (@example-jsmith)

# first set
a: b@c d
e: b@c d

 # second set
f: g@h
i: g@h j
`)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("-------------\nline-based:\n%v\n", man)

	// Output:
	// -------------
	// 2822:
	// Maintainers: InfoSiftr <github@infosiftr.com> (@infosiftr), Johan Euphrosine <proppy@google.com> (@proppy)
	// SharedTags: latest
	// GitRepo: https://github.com/docker-library/golang.git
	// 
	// Tags: 1.6.1, 1.6, 1
	// GitCommit: 0ce80411b9f41e9c3a21fc0a1bffba6ae761825a
	// Directory: 1.6
	// 
	// Tags: 1.5.3, 1.5
	// SharedTags: latest, 1.5.3-debian, 1.5-debian
	// GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
	// Directory: 1.5
	// 
	// Shared Tag Groups:
	// 
	//   - latest
	//     - 1.6.1, 1.6, 1
	//     - 1.5.3, 1.5
	// 
	//   - 1.5.3-debian, 1.5-debian
	//     - 1.5.3, 1.5
	// 
	// -------------
	// line-based:
	// Maintainers: InfoSiftr <github@infosiftr.com> (@infosiftr), John Smith <jsmith@example.com> (@example-jsmith)
	// GitFetch: refs/heads/*
	// 
	// Tags: a, e
	// GitRepo: b
	// GitCommit: c
	// Directory: d
	// 
	// Tags: f
	// GitRepo: g
	// GitFetch: refs/tags/h
	// GitCommit: FETCH_HEAD
	// 
	// Tags: i
	// GitRepo: g
	// GitFetch: refs/tags/h
	// GitCommit: FETCH_HEAD
	// Directory: j
}