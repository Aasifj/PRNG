# Pseudo-Random Number Generator(PRNG)
##This method of generating random numbers involves computational algorithms that can produce apparently random results.
#Why Pseudo-Random? 
##Because the end results obtained are in fact completely determined by an initial value also known as the seed value or key. Therefore, if you knew the key value and ##how the algorithm works, you could reproduce these seemingly random results.

#The Linear Congruential Generator
This generator produces a series of pseudorandom numbers. Given an initial seed X0 and integer parameters a as the multiplier, b as the increment, and m as the modulus, the generator is defined by the linear relation: Xn ≡ (aXn-1 + b)mod n.

##Reference: https://en.wikipedia.org/wiki/Linear_congruential_generator

