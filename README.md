# sachista-chess

## GO lang version

2021-08

https://gitlab.com/dsaiko/sachista-chess-go

New rewrite of the engine into Go language.

I like Go. It is kind of safe C with garbage collector for me. The quality and readability of code of the latest
implementation should be better than all previous versions, ease of build and distribution chain is amazing.

Still, the C++ version (with multithreading) is unbeaten in terms of performance.

### Latest performance results:

2021-08

**AMD Ryzen 7 3700X - Windows 10 WSL / Ubuntu**

    sachista-chess-go perfT finished:
    FEN:    rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
    depth:  7
    count:  3,195,901,860
    time:   6.5251358s

    sachista-chess (c++ version) PERFT finished:
    FEN  : rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
    depth: 7
    count: 3,195,901,860
    time : 2 [seconds]
    speed: 1,292,840,558 [/second]

**MacBook Pro (Retina, 15-inch, Mid 2015)**

Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz

    sachista-chess-go perfT finished:
    FEN:    rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
    depth:  7
    count:  3,195,901,860
    time:   13.728322433s

    sachista-chess (c++ version) PERFT finished:
    FEN  : rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
    depth: 7
    count: 3,195,901,860
    time : 4 [seconds]
    speed: 655,838,674 [/second]

### Swift Version

https://gitlab.com/dsaiko/sachista-chess-swift

Swift version of the engine

I was not able to make it fast enough due to lack of support for native arrays.

## C++ Version 0.2.0

2015-03

https://gitlab.com/dsaiko/sachista-chess

Move generator rewritten from C into C++. Code quality is better, can now be compiled using Microsoft compiler, uses
more functionality from std c++11 library (strings, streams, chrono, collections, threads ...). No More OpenMP.

* Tested on:
  * MS VC++ 2013 x86/x64
  * MinGW 4.9 x86/x64
  * gcc 4.9 x32/x64
  * CLang 3.6.0 x32/x64

Build system was changed from cmake to qmake (without using any of Qt libs). Reason: better handling of qmake projects
inside QtCreator (compared to cmake projects), mainly switching the projects target compilers and release / debug
versions.

Perft results are quite impressive,
see [docs/perft-results.txt](https://gitlab.com/dsaiko/sachista-chess/blob/master/docs/perft-results.txt) for details.

* In short:
  * 30 000 000 per second with copy/make/validate logic
  * cca 80 000 000 per second without making end leave move
  * cca 600 000 000 per second using zobrist hash cache
  * up to 1 200 000 000 per second adding multi threading (on 4 cores)

I was also comparing make/unmake vs copy/make for the chessboard perf test, in contrast to most information at the net,
I have found copy/make slightly faster.

* Best perft 8 results: i5-4690s (DDR3 16GB 2100 MHz), Linux 3.18.6-1-ARCH x64 CLang 3.6.0
  * depth: 8
  * count: 84,998,978,956
  * time : 66 [seconds]
  * speed: 1,287,551,183 [/second]

Voilà! :-)

It was a good excercise to start working again on the engine itself, for now I consider the move generation part to be
completed.

## 2014:

Simple and effective chess move generator and possible future chess engine implementation in C

Sachista-chess is a simple and (hopefully) effective implementation of chess move generator.

It may grow into UCI chess engine in a near future, but so far I just wanted to focus on implementing move-generator in
a way which would satisfy me with both performance and code quality.

This is my probably fourth or fifth rewrite of move generation engine.

What I remember, my first try was quite a long ago, when I implemented chess-moves generation in Java from scratch, then
rewriting that into C++. Without reading any chess-programming articles, this code worked well, but the first move
generation produced only around 20.000 legal moves per a second
(on common hardware relevant to let's say year 2006). Later on, I experimented with rewriting this generator to Scala,
which ... well ... was not really any faster ;-). Just an experiment. I also created a simple Java Swing GUI for my
chess programming that time, I think it looked quite OK. But well, Swing is dead.

In year 2013, I created a "proper" version of a generator and basic UCI chess-engine in Java, implementing
magic-chessboard mechanism. As I really did not like the engine part, I rather deleted it and tried to make just the
generator as good as possible. My code is generating pseudo-legal moves which needs a makemoverification for perft, in
Java version, this was generating around 12.000.000 legal combinations per second in one thread and around 72.000.000
when running in multithreaded version on my current notebook using Java7 fork-and-join framework.

After being upset with Java (will not list the reasons here, but as a long time Java developer I have my own opinion), I
decided to rewrite the engine again, this time into C and hunt performance. I was able to achieve 35.000.000 of legal
moves per a second with my perft(3x faster than almost (90%) the same solution in Java) (still the same notebook), but
comparing those results to stockfish chess engine, I was not happy. Stockfish perft runs around 120.000.000 positions
per seconds, but that's only because it generates legal moves
(vs my pseudo-legal generator) and does not make the move for the last generated ply (no need to make-unmake). Forcing
stockfish to generate even the leaves moves, my performance was better, as stockfish was generating around 25.000.000
legal positions /s.

Adding OpenMP Multithreading and perft caching boosted my perft results to the sky ... see peft results

* Generator in a nutshell:
  * C language, C99 standard, CMake
  * magic bitboards
  * zobrist keys
  * tested on x64 and x86 Linux and Windows platforms, compiled by GCC and ICC
  * OpenMP multi threading
  * CppUTest unit testing suite
  * valid perft with multithreading and position cache
  * no unmake-move, using board-copy for perft
  * posible future paths: chess gui or chess engine

Would like to give a credit to following sources which were my (technical or motivation) inspiration:

* Mediocre Chess blog
* Chess programming wiki
* Stockfish
* Phalanx

note on unmake move: I have implemented unmake move, but that code is not very clear and it still required additional
arrays where non-recoverable info would be stored. Then I deleted unmake move and replaced it simply with currentBoard=*
boardBackup; board backuping. Code is cleaner, performance the same in my case.
