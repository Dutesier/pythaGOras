# pythaGOras 
Recreating the Dining Philosophers Problem... in GO!

## Description
This was done for me to learn a bit more about GO.
It focuses mainly on go routines and mutexes. I've also been testing some interface based architecture (yay me).

## How it works
The program should take the following arguments: 
 - number_of_philosophers time_to_die time_to_eat time_to_sleep [number_of_times_each_philosopher_must_eat]
- number_of_philosophers: is the number of philosophers and also the number of forks.
- time_to_die: is in milliseconds, if a philosopher doesn’t start eating ’time_to_die’ milliseconds after starting their last meal or the beginning of the simulation, it dies.
- time_to_eat: is in milliseconds and is the time it takes for a philosopher to eat. During that time they will need to keep the two forks.
- time_to_sleep: is in milliseconds and is the time the philosopher will spend sleeping.
- number_of_times_each_philosopher_must_eat: argument is optional, if all philosophers eat at least ’number_of_times_each_philosopher_must_eat’ the simulation will stop. If not specified, the simulation will stop only at the death of a philosopher.
- Each philosopher should be given a number from 0 to ’number_of_philosophers - 1’.
- Philosopher number 0 is next to philosopher number ’number_of_philosophers - 1’.

Any other philosopher with the number N is seated between philosopher N - 1 and
philosopher N + 1.
## Program logs
Any change of status of a philosopher must be written as follows (with X replaced with the philosopher number and timestamp_in_ms the current timestamp in milliseconds):
 - * timestamp_in_ms X has taken a fork
 - * timestamp_in_ms X is eating
 - * timestamp_in_ms X is sleeping
 - * timestamp_in_ms X is thinking
 - * timestamp_in_ms X died
The status printed should not be scrambled or intertwined with another philosopher’s status.
We can’t have more than 10 ms between the death of a philosopher and when it will print its death.

Obviously, philosophers should avoid dying!
