bornInUk(a).
bornedAfterCommencement(a).
father(b, a).
isBritish(b).
 
parent(X, Y) :-
 father(X, Y);
 mother(X, Y).
 
britishCitizen(X) :-
 	(
		% A person born in the United Kingdom after commencemen
	 (bornInUk(X),bornedAfterCommencement(X));
		%in a qualifying territory on or after the appointed day
	 (bornInQualifyingTerritory(X),bornedAfterAppointedDay(X))
	 ),
 
	% shall be a British citizen if at the time of the birth his father or mother is—a British citizen
% or settled in the United Kingdom [F2or that territory
	 (
	 parent(Y,X), 
		 (
		 isBritish(Y);
		 isSettledInUK(Y);
		 (
			 bornInQualifyingTerritory(X),
			 isSettledInQualifyingTerritory(Y))
 		)
	 ).
