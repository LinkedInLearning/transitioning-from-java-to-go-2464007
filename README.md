# Transitioning from Java to Go
This is the repository for the LinkedIn Learning course Transitioning from Java to Go. The full course is available from [LinkedIn Learning][lil-course-url].

![Transitioning from Java to Go][lil-thumbnail-url] 

If you are used to working in Java but need to learn to code efficiently in Go, instructor Adelina Simion has you covered. In this course, she shows you how to learn Go quickly and continue delivering code at your regular speed in production. Adelina walks you through basics like program structure, running programs, declaring variables, basic data types, and pointers. She defines functions and walks you through how to use several of them in Go. Adelina goes over structures (structs) and interfaces, including call methods, modules, and interfaces. She explains data structures, such as arrays, slices, maps, loops, and ranges, then describes how to handle JSON. Adelina discusses sorting and searching, as well as CLIs. She steps through how to use unit testing in Go, then finishes up with a discussion of concurrency and generics that introduces channels and goroutines.

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
3. Run `go get .` from the root directory where you cloned this repository to install dependencies used in your current branch. Note, different branches have different dependencies, so make sure to run this command when you change branch.


### Instructor

Adelina Simion 
                                                   

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/adelina-simion).

[lil-course-url]: https://www.linkedin.com/learning/transitioning-from-java-to-go
[lil-thumbnail-url]: https://cdn.lynda.com/course/2464007/2464007-1654622641668-16x9.jpg
