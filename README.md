# Create you a friend with this package
This is the github page of the *twitterfriends* package

If you ever feel alone, don't anymore. Create you a friend with this package and find the hapiness :D

## Just a word

This package isn't organized, nor finished.
It still needs a lot of works and I'm going to do it.
Any suggestions are welcomed :simple_smile:

## For now
There isn't tests for this package, there're coming soon :D

You can still use the package. Here are some steps:

1. Go to the [developer twitter site](https://developer.twitter.com/)
   - Here create an account for your friend
   - Create your friend an app [here](https://developer.twitter.com/en/apps)
   - Go inside the *keys & tokens* section
   - Write down the api key and api secret
   - Then create tokens and write them down
2. Now that we have your friend keys and tokens, his application and account, let's create a **.env** file
   - Create a new file call "app.env" and put it inside the git cloned folder
   - Then, open the file, time for writing some lines
        ```
        CONSUMER_KEY="your friend's api key"
        CONSUMER_SECRET="your friend's api secret"
        TOKEN_KEY="your friend's token"
        TOKEN_SECRET="your friend's token secret"
        ```
3. Everything should works, time to create your main.go file
    - Here is the code for your main.go file (There isn't much more to do for now...)
        ```golang
        package main
        
        import (
            tw "github.com/Betelgeuse1/RepettiRepetto/twitterfriends"
        )
        
        func main() {
            tw.Debug = true // Enable the debugging, still in progress, not necessary
            tw.SendTweet("here is the tweet status, a.k.a the text content of your tweet")
            
        }
        ```
    - Just have to run the main.go file now, and your friend will speak for the first time
        `go build -o myAwesomeFriend; ./myAwesomeFriend`
        
Here you go ! Your friend talks for the first time, congrats !

## What can it do ?
#### SendTweet(*status string*) *func*
Send a tweet with the status as the text content of it

#### Debug *boolean*
Enable/Disable the request and response to be consoled-printed