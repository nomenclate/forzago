forzago

Summary 

Forzao is a Basic utility and library for dealing with ForzaDataOut from Forza 7 Motorsport.  This was largely written as a learning exercise.  code  from Jean de Klerk's scalability talk https://github.com/jadekler/git-go-scalability-talk was used as a starting point and then added to from there.

Info on forzadataout can be found in the following forum post:

https://forums.forzamotorsport.net/turn10_postsm926839_Forza-Motorsport-7--Data-Out--feature-details.aspx#post_926839


There are three main components in forzago, inputters, outputters, and queues.  

Inputter place a []bytes represeenting a single forza dataout back onto the queue to be later removed and dealt with as appropriate by an Outputter.

Forzago was intended to allow for mixing and match of differnt inputters (i.e. from network, from file etc.) and outputters (i.d. StdOut, Json, influx line protocol). 

Interfaces are provided for interacting with as well as writing additonal variantions of inputters/outputters/queues
