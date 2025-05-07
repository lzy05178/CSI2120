even(0):-!.
odd(N):-N>0,
    M is N-1,
    even(M).
even(N) :-N>0,
    M is N-1,
    odd(M ).

sum_odd_numbers(List,Sum):-
    findall(X,(member(X,List), odd(X)),Result),
    sum(Result,Sum).

sum([],0).

sum([X|T],Sum):-
    sum_odd_numbers(T,Left),
    Sum is X+Left.


