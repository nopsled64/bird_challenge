# Solutions
Spoiler warning! Don't read this unless you want the answers spoiled for you!

These answers do not cover the bonus section.

## Question 1
Data can be represented as a series of 1s and 0s, also known as binary. If this is true, then data can also be represented as Cheeps and Chirps! If you replace all of the Chirps with 0s, and replace all Cheeps with 1s, you get a sequence of binary data:

```
00101111011000110110100001100001011001100110011001101001011011100110001101101000
```

If you then convert this from Binary to Text using any kind of converter available online (I suggest CyberChef!), you get the next location `/chaffinch`

## Question 2
This question involves "decrypting" an MD5 hash. MD5 is a hashing algorithm which can convert any kind of data into a 32 character sequence. It is designed as a one-way function, so you can easily convert data from its original state into this hash, but is theoretically impossible to convert it back to its original state.

The conversion is always going to be the same, so if I were to create an MD5 hash out of the word "fox" it will always be `2b95d1f09b8b66c5c43622a4d9ec9a04`. Because of this there are many services that have made large lookup tables of words against hashes. So with a simple Google search for `b6bc6539453d2f0722db8cc3ac5730a2` we can see the phrase `sparrow`.

MD5 hashes are sometimes used to encrypt passwords, and this challenge hopefully makes it obvious why that's a terrible idea!

## Question 3
This one is a lot more involved. The page mentions that we're working with Base64, which is another way of representing data. By using a simple online base64 to text converter (or CyberChef, which I highly recommend!), we get the text output:
```bash
#!/bin/bash
Seagull="Eea2af6C7Vc+AHE98VJZwH6RIAAAA=";Baygull="H4sIAAToMF0A/wVAMQ0AMAyygoM5qYuG7I";Peacock="a";Pigeon="e";Pelican="w";Parrot="f";Partridge="o";Parakeet="g";Penguin="c";Puffin="r";Pipipi="i";Pyrrhuloxia="p";This=$Parrot$Pelican$Partridge$Puffin$Penguin$Parakeet$Pyrrhuloxia;echo -n "The answer to question 3 is: ";echo ${This:1:4}|rev;openssl enc -base64 -d <<< ${Baygull}${Seagull}|gunzip -cf|rev;echo " ";
```
Looking at the very start of this block of text we see the phrase `#!/bin/bash`, this is an indicator that it is a shell script designed to be run with BASH. This is also the default shell for most Linux distributions, so this part of the task is very easy if you have access to a Linux system! If you do, you can simply run the script and get the answer to this question.

Otherwise, it is also possible to work out what the script is doing by reading it line by line, but this might take a while as I've tried to make it a little hard to read!

The script has two outputs, the answer to question 3, and the location of the next set of questions: `/penguin`

## Question 4
The answer for this one is given immediately upon visiting the page, so I won't explain it here!

## Question 5
This question sends you to `/bird.mp3`, which is simply a recording of me saying `Eagle` in reverse!

## Question 6
This one is a simple general knowledge / Googling skills question. The musical group Kid Creole & The Coconuts had a song called `Stool Pigeon`, which is a slang term for a police informer / decoy.

## Question 7
After question 6 is a section that mentions the rest of the questions are at `\x2f\x70\x75\x66\x66\x69\x6e`. This is Hexadecimal, which is yet another way of representing data. after decoding this using CyberChef (other decoders are available) we get `/puffin`.

The top of the page asks "what is the bird at `/bird.jpg`?", and by navigating to it we get a nice picture of a bird. If you're a bird expert you might be able to identify it straight away, but if you're a regular bird-blind person like me you can do a reverse image search using something like Google or TinEye. You might have encountered problems searching based on the URL as I've restricted access to the bird-shirt site.

After doing the reverse image search, you can see it is a `Eurasian Jay` (though I'll also accept just `Jay`)

## Question 8
Here we're asked who the human is at `/human.jpg`, which is an image of the director Brad _Bird_.

It appears most reverse image searches won't give you this straight away, so requires some deeper searching. Some clues include the badge he is wearing on his lapel, which is the Incredibles logo, one of the films he has directed.

## Question 9
This is a different one! It mentions the answer is on Slack.

**Note:** This question requires some extra setup, make sure that the solution is on Slack or change this question.

## Question 10
The final question requires a bit of research into endangered birds, but on the site is a PDF report that lists all of the UK Birds of Conservation Concern. Among them is the `White-tailed eagle`