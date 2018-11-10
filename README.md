# gorocksim

This simulates (in 1D) the upward trajectory of a model rocket, you enter the particulars of a rocket, it reads the profiles of model rocket engines (which gives you thrust and mass at time t), and with some fancy math you can get instanteneous velocity and height.

(note to run this you have to move the contents of the engines directory to ~/.gorocksim/engines or else they cant be found....)

# components of this app

# the parser
this app can parse either rse files (see thrustcurve.org) or wrasp files for the thrust profiles of model rocket engines.   Wrasp files are kinda yesterdays news because they arent as detailled.  RSE files give you much better info for simulations (ideally they would be the same).

** this part is broken right now **

# the repository

stores information on rockets, mass, diameter, etc


# the simulation

this part starts at time zero and steps through the thrust profile of the selected engine data and using Runge-Kutta formulas simulates the instantaneous thrust at time t.   This is only a one dimensional simulation (I.e. up and down) and doesnt consider motion in other than the y direction.  Its rough but its somewhat accurate.   Real world measurements with the algorithm show it to be correct within 10%.





