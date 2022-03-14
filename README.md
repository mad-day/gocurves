# gocurves
Elliptic Curves for go's `crypto/elliptic` package.

[![GoDoc](https://godocs.io/github.com/mad-day/gocurves?status.svg)](https://godoc.org/github.com/mad-day/gocurves)

Additional Elliptic Curves for `crypto/elliptic`.
Curves are implemented using `*elliptic.CurveParams`.

## About the Brainpool Standard Curves [RFC5639].

The ECC Brainpool, is a working group of the state-industrial association TeleTrusT
(members including BKA, BSI) on the subject of Elliptic Curve Cryptography.
The working group specified a number of elliptic curves in 2005, which were standardized
in March 2010 in RFC 5639 of the IETF.

For these curves, the choice of bit length 512 should be mentioned, in contrast to the bit
length 521 preferred by many other institutions (e.g. NIST, SECG).

POSSIBLE KLEPTOGRAPHY: The design space of the Brainpool curves contains so many degrees of
freedom that a back door cannot be excluded with certainty. The Brainpool curves are also
uncertain about some desirable properties.

Of the Brainpool curves only the brainpoolP*t1 variants are implemented because of
crypto/elliptic's limitation, that the domain parameter `A` must be `-3`.

## About the Microsoft's Nothing Up My Sleeve (NUMS) curves.

These curves are elliptic curves over a prime field, just like the NIST or Brainpool curves.
However, the domain-parameters are choosen using a VERY TIGHT DESIGN SPACE to ensure, that
the introduction of a backdoor is infeasable. For a desired size of `s` bits the prime `p` is
choosen as `p = 2^s - c` with the smallest `c` where `c>0` and `p mod 4 = 3` and `p` being prime.

For Weierstrass curves (with `a = -3` for backward compatibility) `b` is choosen the smallest
`abs(b)` where the prime order `n = #Eb(GF(p))` of the curve is as high as possible, dealing
a highter `n` for a larger `abs(b)`.

See [https://tools.ietf.org/html/draft-black-numscurves-02](https://tools.ietf.org/html/draft-black-numscurves-02) and [https://eprint.iacr.org/2014/130](https://eprint.iacr.org/2014/130)

Of the NUMS curves only the Weierstrass curves are implemented.

