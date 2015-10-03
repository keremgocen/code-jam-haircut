# code-jam-haircut

Solution to the Haircut problem on Google Code Jam, in golang.

Source: https://code.google.com/codejam/contest/4224486/dashboard#s=p1

## Algorithm

First, determine how long it would take for the barbers to cut the hair of every
person in front of us - this gives us the time at which we are guaranteed to be
getting our hair cut. We do this by determining a lower (0 minutes) and upper
bound in time for which cutting all their hair must take, then binary searching
solutions until we find the minute at which we would be attended.

We know based on the time A. which barbers are available and B. how many people
have been served. We assign people to open barbers in our given minute until we
find our barber.
