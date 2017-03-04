### What does it do?

This app converts input text file that looks like this

          99999
         9888788889
         5551141
        5251114111
        52551115111
         511115555
          111111

into a repository with commits, which contribution activity later looks as follows

![Mario's head](./models/mario_head.png)

### How to run?
First run `make` to download all dependencies and run tests. After that, execute:

    go run ./src/app/*App.go -offset 20 -input-model ./models/mario_head.txt -output-repo ./output_repo -git-config ~/.gitconfig
    
To print help message:

    go run ./src/app/*App.go -help
    
### Why gophers?

Why go, I hear you ask. I had this small project idea and to keep it interesting I had to somehow increase difficulty level. 
I had never used golang before, so it was good opportunity to see what it is all about. As a result, the application may be
in few places avoidably complicated (e.g. usage of channels and go routines). 
