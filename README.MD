# msds301-wk8

### Setup (windows, gitbash)
- Clone repo with `git clone git@github.com:jeremycruzz/msds301-wk8.git`
- Make sure you have `g++` and `MYSS2`[instructions here](https://code.visualstudio.com/docs/cpp/config-mingw)
    - This is for the sqlite driver
- Run `go mod tidy`

### Building executable
- Run `go build ./cmd/summarizer`

### Testing
- Run `go test ./...`
- Right now only SummarizerService is tested.

### Running Go executable
- Run `./summarizer.exe -apikey {API_KEY} -p :8080` 
    - `-p` flag is optional `:8080` is default
    - I'll put the apikey as a comment in the assignment

- Endpoints
    - `GET  localhost:8080/api/v1/summarize/{topic}`
        - Gets a topic from db or generates from chatgpt and then stores it in the database
        - Topic is actually the title of a wikipedia article for now
    - `POST localhost:8080/api/v1/summarize/{topic}/update`
        - Updates a an existing topic with a new response from chatgpt
- Html page
    - `localhost:8080/api/v1/summarize`
    - Its kinda ugly.
- Notes
    - Some wikipedia pages might be too long to send to chatgpt so try to stick to shorter wiki pages.
    - \$0.82 / \$10.00 has been used so don't go crazy!
    - Existing rows are
      - hero
      - qt (software)
      - robot
      - tea
### Background / Conclusion

![Design doc](/assets/design.png)
For the project I made a hybrid between a selection tool and AI-assisted information extraction (kind of). The whole point of this was to have chatgpt proccess some text in two different ways and display it to the user. Right now I'm using the wikipedia api over the scraper from the previous project due to ease of use and the api removed the html for me. This doesn't have to be wikipedia. It could be company docs, different product pages, code, etc.
<details>

<summary> Flowchart For Summarize </summary>

![flowchart](/assets/flowchart.png)

</details>

For the Db I decided to use sqlite so I didn't have to deal with docker or hosting a db somewhere. Ideally I would probably use MongoDb or another NoSql Db so that I can store multiple responses from chat gpt and allow the user to scroll through them on the page. Other than the Db there are a few improvements I want to work on in the future:

  - Better queries
    - Not require exact wikipedia title

  - Better error handling

  - Better UI
    - Let the user know that the page is loading
    - Format the summaries better
    - Allow the UI to refresh responses
    - Probably use React or experiment with svelte
  
  - Testing
    - Test everything else

  - Customize Prompts

With adjustments, a tool like this would be very useful for all sorts of professions. As a software engineer, I could use this tool to translate product/marketing/finance docs into documents that I could understand or even translate my docs so that others could easily understand them.
