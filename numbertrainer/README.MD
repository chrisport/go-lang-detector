# Chinese number trainer
Learn Chinese numbers `27   èr shí qī    二十七`.   
## Usage

Simply use ```go run chinese.go```

You can chose one of 3 modes:
- **Active**: You will see a random number 1-99, write down Chinese or Pinyin youself, press <Enter> to see if you were correct. For example 27
- **Passive Chinese letter**: You will see a number in Chinese letters, type the correct number and press <Enter> to check. e.g. `三 --> 3 <Enter>`
- **Passive Pinyin**: You will see a random number 1-99 in Pinyin, write the correct number and press <Enter> to check. e.g. `sān --> 3 <Enter>
`

It will look like this:
```
Chose mode:
  1 ) Active
  2 ) Passive Chinese letters
  3 ) Passive Pinyin
$ 2

Please answer in numbers and confirm with Enter, e.g. 三 --> 3 <Enter>
______________________________
ROUND 1
七十八

$ 78

78    qī shí bā      七十八
CORRECT
______________________________
ROUND 2
六

$ 6

6     liù            六
CORRECT
______________________________
ROUND 3
六十一

$ 31

61    liù shí yī     六十一
WRONG
______________________________
```
