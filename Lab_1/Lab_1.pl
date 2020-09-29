parent(gaius_Julius_Caesar_the_Elder, gaius_Julius_Caesar).
parent(aurelia_Cotta, gaius_Julius_Caesar).

parent(gaius_Julius_Caesar_the_Elder, julia_Caesar_Major).
parent(aurelia_Cotta, julia_Caesar_Major).

parent(gaius_Julius_Caesar_the_Elder, julia_Caesar_Minor).
parent(aurelia_Cotta, julia_Caesar_Minor).

parent(gaius_Julius_Caesar, julia_Caesar).
parent(cornelia_Cinilla, julia_Caesar).

parent(marcus_Atius_Balbus, atia_Balba_Caesonia).
parent(julia_Caesar_Minor, atia_Balba_Caesonia).

parent(gaius_Octavius, gaius_Octavaius_Augustus).
parent(atia_Balba_Caesonia, gaius_Octavaius_Augustus).

parent(gaius_Octavius, octavia_Minor).
parent(atia_Balba_Caesonia, octavia_Minor).

parent(marcus_Antonius, antonia_the_Elder).
parent(octavia_Minor, antonia_the_Elder).

parent(gaius_Octavaius_Augustus, julia_the_Elder).
parent(cloudia_Pulhra, julia_the_Elder).

parent(gaius_Octavaius_Augustus, tiberius_Caesar_August).
parent(livia_Druzilla, tiberius_Caesar_August).

parent(tiberius_Claudius_Neron_the_Elder, neron_Claudius_Druz_Germanic).
parent(livia_Druzilla, neron_Claudius_Druz_Germanic).


adoptive_parent(gaius_Julius_Caesar, gaius_Octavaius_Augustus).
adoptive_parent(calpurnia_Pizonis, gaius_Octavaius_Augustus).


spouse(gaius_Julius_Caesar_the_Elder, aurelia_Cotta).
spouse(gaius_Julius_Caesar, cornelia_Cinilla).
spouse(gaius_Julius_Caesar, calpurnia_Pizonis).
spouse(marcus_Atius_Balbus, julia_Caesar_Minor).
spouse(gaius_Octavius, atia_Balba_Caesonia).
spouse(marcus_Antonius, octavia_Minor).
spouse(gaius_Octavaius_Augustus, cloudia_Pulhra).
spouse(gaius_Octavaius_Augustus, livia_Druzilla).
spouse(gaius_Octavaius_Augustus, scriboniya_Libo).
spouse(tiberius_Claudius_Neron_the_Elder, livia_Druzilla).

male(gaius_Julius_Caesar_the_Elder).
male(gaius_Julius_Caesar).
male(marcus_Atius_Balbus).
male(gaius_Octavius).
male(marcus_Antonius).
male(gaius_Octavaius_Augustus).
male(tiberius_Caesar_August).
male(neron_Claudius_Druz_Germanic).
male(tiberius_Claudius_Neron_the_Elder).

female(aurelia_Cotta).
female(cornelia_Cinilla).
female(julia_Caesar_Minor).
female(atia_Balba_Caesonia).
female(octavia_Minor).
female(cloudia_Pulhra).
female(livia_Druzilla).
female(julia_Caesar_Major).
female(julia_Caesar_Minor).
female(julia_Caesar).
female(antonia_the_Elder).
female(julia_the_Elder).
female(calpurnia_Pizonis).
female(scriboniya_Libo).


grandparent(X,Y):- parent(X, Z), parent(Z, Y).
adoptive_grandparent(X,Y):- parent(X, Z), adoptive_parent(Z, Y).
brother(X,Y):- parent(Z, X), parent(Z,Y), X\==Y, male(X).
sister(X,Y):- parent(Z, X), parent(Z,Y), X\==Y, female(X).
uncle(X,Y):- parent(Z, Y), brother(X, Z).
aunt(X,Y):- parent(Z, Y), sister(X, Z).
predecessor(X,Y):- parent(X,Y).
predecessor(X,Y):- parent(X,Z), predecessor(Z,Y).
mother(X,Y):- parent(X,Y), female(X).
father(X,Y):- parent(X,Y), male(X).
husband(X,Y):- spouse(X,Y), male(X).
wife(X,Y):- spouse(X,Y), female(Y).