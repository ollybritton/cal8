# cal8
`cal8` is an experiment in a different calendar system.

```
      July 2020 
Mo Tu We Th Fr Sa Su Va 
       1  2  3  4  5  6
 7  8  9 10 11 12 13 14
15 16 17 18 19 20 21 22
23 24 25 26 27 28 29 30
31
```

Under the `cal8` system, weeks are 8 days long consisting of Monday-Sunday plus an extra "Valday". During lockdown, for me time has been meaningless since every day has been pretty much the same (I really don't get out much). I thought it would be cool to try out a different system of keeping time since it's a good opportunity to do so.

For me, having an extra day gives two main benefits:

* Extra running rest day. I've been following [this](https://jgorunning.wordpress.com/2010/06/03/marathon-training-schedule-template-v1-0/) marathon training plan for a while but find running 6/7 days a week difficult. Having an extra rest day on *Valday* eases the pressure a little bit.
* More time to recover from cheat day. I've been losely following the Slow Carb diet which means a cheat day once a week. The only problem is that the sheer amount of food I eat on Friday means that I'm basically back to square one for that week. Having an extra day to recover would be nice.

### Installation
If you're installing this you're probably crazy.

```
go get -u github.com/ollybritton/cal8/...
```

### Usage
The `cal8` command aims to provide an set of functionality similar to that of the Unix `cal` utility.

```sh
cal8 # Prints the current month's calendar.
cal8 2020 # Prints the year's calendar.
cal8 -3 # Prints a 3 month calendar.
cal8 --months/-n 10 # Displays a 10 month calendar.

cal8 next january # Shows calendar for next january, relative time parsing thanks to https://github.com/tj/go-naturaldate
```

The command also comes with a `when` subcommand, which is useful for converting normal dates into the new calendar system. For example

```sh
cal8 when Saturday
# Output: 2020-07-11 Saturday

cal8 when 2020-07-15
# Output: 2020-07-15 Tuesday (Wednesday non-cal8)
```

### Maths
The calendar is calculated using a modified version of [Zeller's Algorithm](https://en.wikipedia.org/wiki/Zeller%27s_congruence).