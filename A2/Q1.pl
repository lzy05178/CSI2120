parent(john, mary).
parent(john, tom).
parent(mary, ann).
parent(mary, fred).
parent(tom, liz).
male(john).
male(tom).
male(fred).

female(mary).
female(ann).
female(liz).



parent(Parent, Child) :-parent(Parent, Child),!.

checkparent(Parent, Child) :-parent(Parent, Child),!.

sibling(Sibling1, Sibling2):-
    parent(X,Sibling1),
    parent(X,Sibling2),
    Sibling1/=Sibling2.

grandparent(Grandparent, Grandchild):-
    parent(X,Grandchild),
    parent(Grandparent, X).

ancestor(Ancestor, Descendant):-
    parent(Ancestor, Descendant);
ancestor(Ancestor, Descendant):-
    parent(X, Descendant),
    ancestor(Ancestor, X).