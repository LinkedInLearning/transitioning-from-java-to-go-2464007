# Transitioning from Java to Go
This is the repository for the LinkedIn Learning course `Transitioning from Java to Go`. The full course is available from [LinkedIn Learning][lil-course-url].

_See the readme file in the main branch for updated instructions and information._
## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Go development tools](https://go.dev/doc/install)
    - [Visual Studio Code](https://code.visualstudio.com/) or any other IDE of your choice
    - [Git](https://git-scm.com/)
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.
3. Run `go install github.com/golang/mock/mockgen@v1.6.0` to install `mockgen`.
4. Add `mockgen` install location to your `PATH` environment variable using `export PATH=“$HOME/go/bin:$PATH"`
5. Run `go get github.com/golang/mock/` to install `golang/mock`.
6. Run `go get .` from the root directory where you cloned this repository to install dependencies used in your current branch. Note, different branches have different dependencies, so make sure to run this command when you change branch.


[0]: # (Replace these placeholder URLs with actual course URLs)

[lil-course-url]: https://www.linkedin.com/learning/
[lil-thumbnail-url]: http://

