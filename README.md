### What does it do?

This app converts input text file that looks like this

          99999
         9888788889
         5551141
        5251114111
        52551115111
         511115555
          111111

into a repository with commits, which GitHub contribution activity later looks as follows

![Mario's head](./models/mario_head.png)

### How to run it?
First run `make` to download all dependencies and run tests. After that, execute:

    go run ./src/app/*App.go -offset 20 -input-model ./models/mario_head.txt -output-repo ./output_repo -git-config ~/.gitconfig
    
To print help message:

    go run ./src/app/*App.go -help
    
### Why gophers?

Why Go, I hear you ask. I had this small project idea and needed to keep it interesting. I had never used golang before, 
so it was a good opportunity to see what it is all about. As a result, in few places the application may be avoidably 
complicated (e.g. usage of channels and go routines).