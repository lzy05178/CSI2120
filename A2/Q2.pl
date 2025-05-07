pet(fido, dog, 3).
pet(spot, dog, 5).
pet(mittens, cat, 2).
pet(tweety, bird, 1).
male(fido).
male(spot).
female(mittens).

pet(Name, Species, Age):-pet(Name, Species, Age),!.
checkpet(Name, Species, Age):-pet(Name, Species, Age),!.


species(Species,Count):-
    findall(Nname,pet(Nname,Species,_),X),
    length(X, Count).

age_range(MinAge, MaxAge, Count) :-
    findall(Name,
            (pet(Name,_,Age),Age>=MinAge,Age =<MaxAge),
            X),
     length(X, Count).