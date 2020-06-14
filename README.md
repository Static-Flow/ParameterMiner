# ParameterMiner
Built on a lazy Sunday after seeing this tweet (https://twitter.com/intigriti/status/1272145863868104705?s=20) I present to you, ParameterMiner! Pipe in a list of javascript urls and ParameterMiner pulls all the variable names.  


#### USAGE:

````
  -length int
        Minimum length variable to collect
  -s    saves output to file named from base url of source
  -save
        saves output to file named from base url of source
````

#### Example with single URL:

````echo https://news.ycombinator.com/hn.js?qjUdg9dZheJfdx2Zo5qF | paramMiner.exe````


#### Example output:

````
j
on
n1
rks
a
id
unv
ks
el
sp
n
up
trs
req
url
i
s
pair
next
````

#### Example with single URL and Length Filter:

````echo https://news.ycombinator.com/hn.js?qjUdg9dZheJfdx2Zo5qF | paramMiner.exe -l 1````


#### Example Output with single URL and Length Filter:

````
unv
ks
req
next
id
el
pair
url
up
trs
n1
sp
rks
on
````
