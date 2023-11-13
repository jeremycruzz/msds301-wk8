# msds301-wk7

### Setup (windows, gitbash)
- Clone repo with `git clone git@github.com:jeremycruzz/msds301-wk8.git`
- Make sure you have `g++` and `MYSS2`[instructions here](https://code.visualstudio.com/docs/cpp/config-mingw)
    - This is for the sqlite driver
- Run `go mod tidy`

### Building executable
- Run `go build ./cmd/summarizer`


### Running Go executable
- Run `./summarizer.exe -apikey {API_KEY} -p :8080` 
    - `-p` flag is optional `:8080` is default
    - I'll put the apikey as a comment in the assignment

- Endpoints
    - `GET  localhost:8080/summarize/{topic}`
        - Gets a topic from db or generates from chatgpt and then stores it in the database
        - Topic is actually the title of a wikipedia article for now
    - `POST localhost:8080/summarize/{topic}/update`
        - Updates a an existing topic with a new response from chatgpt

### Background

TODO

### Conclusion

TODO


