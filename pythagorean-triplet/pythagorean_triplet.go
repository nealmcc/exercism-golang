package pythagorean

import (
  "math"
  "sort"
  "strconv"
)

// a Triplet (or Pythagorean Triplet) is a set of three integers,
// a, b, c such that a < b < c and a^2 + b^2 = c^2
type Triplet [3]int

// Range finds all of the pythagorean triplets such that a >= min and c <= max
func Range(min, max int) []Triplet {
  var t []Triplet

  for c := max; c >= min+2; c-- {
    c2 := c * c
    for a, a2 := min, min*min; a2 < c2/2; a, a2 = a+1, (a+1)*(a+1) {
      b2 := c2 - a2
      b := int(math.Sqrt(float64(b2)))
      if b*b == b2 {
        t = append(t, Triplet{a, b, c})
      }
    }
  }

  sort.Sort(tripletSlice(t))
  return t
}

// Sum finds all of the pythagorean triplets such that a + b + c = p
func Sum(p int) []Triplet {
  var trips []Triplet
  min, max := (p/3)+1, (p/2)-1

  for c := max; c >= min; c-- {
    b24ac := c*c + 2*p*c - p*p
    if b24ac <= 0 {
      continue
    }

    sqrt := int(math.Sqrt(float64(b24ac)))
    if sqrt*sqrt != b24ac {
      continue
    }

    a := (p - c - sqrt) / 2
    b := (p - c + sqrt) / 2
    trips = append(trips, Triplet{a, b, c})
  }

  sort.Sort(tripletSlice(trips))
  return trips
}

type tripletSlice []Triplet

var _ sort.Interface = tripletSlice{}

// Len is part of sort.Interface
func (s tripletSlice) Len() int {
  return len(s)
}

// Less is part of sort.Interface
// This method compares Triplets lexicographically
func (s tripletSlice) Less(i, j int) bool {
  for n := 0; n < 2; n++ {
    a := strconv.Itoa(s[i][n])
    b := strconv.Itoa(s[j][n])
    if a < b {
      return true
    }
    if b < a {
      return false
    }
  }

  a := strconv.Itoa(s[i][2])
  b := strconv.Itoa(s[j][2])
  return a < b
}

// Swap is part of sort.Interface
func (s tripletSlice) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

/*
Reasoning:
----
I can tell this has to do with conic sections, but it has been a *long* time
since I did anything with them. So I'm going to think this through using basic
geometry initially, and go from there.

let x = a, y = b, and treat n and c as constants:

by definition:
(1) x² + y² = c²
(2) x + y + c = p
(3) x < y < c

=>

(1a) y² = c² - x²       a circle with radius c
(2a) y  = -1x + (p-c)   a line with slope -1, intersects both axes at p-c


In order to solve both equations, we will find the points of intersection:

set y² = y²

=> c² -x² = (p -c -x)²

=> c² -x² = p² -pc -px -pc +c² +cx -px +cx +x²

in standard form:

(4) 0 = x² +(c-p)x + (½p² -pc)

If there is a valid triplet for the current values of p and c, then it will
be (x1, x2, c) where x1 and x1 are the two positive integer roots of (4)

To solve for those roots, we can use the quadratic formula:
x1, x2 = -b±√(b²-4ac)
          -----------
              2a

where
a = 1
b = c-p
c = ½p² -pc

=>

x1, x2 = - (c-p) ±√[(c-p)² -4(½p² -pc)]
           ----------------------------
                       2

       =  p -c ±√[c² -2pc +p² -2p² +4pc]
          ------------------------------
                      2

       = p -c ±√[c² +2pc -p²]
         --------------------
                 2

Now, in the case of solving for Sum, we're being given p. In order to satisfy
inequality (3), the possible values for c will be: p/3 < c < p/2. Otherwise,
a, b, c won't form a triangle.

This should narrow down our candidates for c sufficiently that we can try each
candidate, and calculate b24ac = (c² +2pc -p²). If this value is > 0 and a
perfect square, then we can continue to find x1 and x2.

By example with n = 12:
----

our valid range for c is :

12       12
-- < c < --
 3        2

as integers:

4 < c < 6

.. so we have one candidate to try: c=5, let's try it:

b24ac = c² +2pc -p²
      = 25 + 2*12*5 - 144
      = 1

sqrt(1) is indeed an integer, so continue

a = (p -c -1)/2 = 3
b = (p -c +1)/2 = 4

We have our triplet:
[3,4,5]

note: for the 'Range' method, I'm just using brute force.
*/
