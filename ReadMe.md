# Graceful Termination
 

The idea of this project would probably be useless for many of you 😅. But I am coding for fun. Though the software could be useless, 
the learning of it taught be many new things about golang programming. 

Anyway let's go dive into the roots of this work. 

I have been suffering and annoyed with address already in use when I was trying to start a connection and 
stopping it with Operating systems signals such as `ctrl + Z` and `ctrl + C`. So I thought of writing a 
simple go wrapper around the commands that I execute that can handle the command execution on kill the process 
upon receieving terminationg signals. 

I have named the wrapper as execute. 

So basically you can put this binary in your path and then run any command with the following structure. 


` ./execute pwd`  for example or a complex command like 

`./execute systemctl start mysql` 

So when you want to terminate the process if it's a running process of a command you execute, you can easily use `ctrl + Z` or `ctrl + C` and it will make sure that the underlying process is killed too... 

I am improving this small wrapper and I would like anyone's comment and ideas on this and any further usage we can make to this. 

Issues are always welcome. 

Cheers... 