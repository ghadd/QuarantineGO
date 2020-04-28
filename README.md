# QuarantineGO

[QuarantineGO](https://quarantinego.online) course stuff

## course started 12.04.2020

### Lection #0 Intro to course // 12.04.2020

It was said about the overall course aim and its rules e.g

* discipline
* programming inevitability in todays world
* prerequisities for course
* grading system
* some others...

### Lection #1 Intro to Go Syntax // 14.04.2020

Falshstart xD from making an X-MAS tree in console. Basically, for me this was truly easy as I'm down to programming quite a long time and I had a python script, which did the same, moreover it was blinking with different colors. So, basically

* WORKSPACE ==> really positive about such kind of a structure. Reminds me some Java Maven project one.
* Variables declarations
* loops
* conditionals
* (fmt => stands for `formatted` and is used for standart input and output)

### Lection #2 Variables and Strings // 16.04.2020

It was a good discussion and overlook of basic golang variables syntax and strings usage (especially with runes and stuff). Main points

* variables
* constants
* inline declarations
* underscore syntax
* ABSENCE OF CHARS --> ps. they are somthing as tight dependence bytes (which is int8 alias for ASCII stuff) and runes (an alias for int32 to represent the whole Unicode (UTF-8 encoding) literals)
* if-else statements
* switch statement
* for loop with multiple declarations in it

*We are getting ready to GO!*

### Lection #3 Q&A // 18.04.2020

The Q&A session started on summing up the week and overall situation among the students.

1. 
> Q. "When is it ok to attend to be hired by IT companies?" 

> A. "ASAP. Really. You gotta have to try attending to the companies (no matter what level of this company is -> start from the small and go higher). From the time when you started studying in the uni and you feel that you have some potential - just go ahead."

2. 
> Q. "What language it's better to choose when you start learing programming stuff?"

> A. "You gotta choose really unpopular language. That will help you not to concentrate the "only wanted" language, but to understand the programming origin."

3. 
> Q. "Do I have to try harder myself or ask someone for help?"

> A. "Definitely DO NOT stop on something on a long period. Try asking people who may know the solution or just skip the problem for a while. The solution is going to come to your mind suddenly."

4. 
> Q. "Do I have to know OOP?"

> A. "Sounds like self-answering question. Certainly you have to. This is the kind of basics every programmer must know no matter if you are using functional languages. Moders problems require modern solutions!"

5. 
> Q. "What university it's better to graduate from?"

> A. "The university doesn't really matter. What matters - is the place it is located. It's better to choose the city, which has a bunch of IT companies located in. That's gotta work out."

### Lection #4 Writing simple algorithms // 21.04.2020

Along the lection, it was discussed how to dive into codind something. Especially for bigger projects, splitting data into pieces and KISS-rule* insructions is really important to prevent messing up whilst coding. It was written a small program to solve the quadratic equation. The cool thing is GO's acceptance of complex numbers in its native package cmplx.

*KISS-rule stands for "Keep is simple, stupid" - the key rule of objectivism.

### Lection #5 Arrays and slices // 23.04.2020

We were talking about arrays and slices as contigous data containers and wrote a simple program that demontrates them in work. So, basically arrays are underlaying the slices, what is you rarely will access arrays in go. Slices by origin are immutable unless you overcome their size. Then you have to perform append to them and at the point where size is overcoming the capacity, everything from the existing array copies to the one, that has twice as much the capacity. The array/slice syntax is pretty simple and multidimensional arrays are welcome. Also we discussed the `make` syntax, so that is is able to allocate memory for arrays/slices. Not a big deal.

### Lection 7 // 28.04.2020

Functions, recursion and hashmaps. Basic syntax and application. Introdusing graphics library [GG](github.com/fogleman/gg) and it's basic functions.