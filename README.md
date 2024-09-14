# go-botley

## What's this?

It's a reverse engineering project for the botley toy. I was intersted how the protocol works and build a small library around for simpler coding. 
## How to write program?

```
program := command.Program{
		Name: "botleyTest",
		Bursts: []command.Burst{
			command.Backward,
			command.Rotate90,
			command.Rotate90,
		},
	}
outputFlipperIrFile("/tmp/botley_synth.ir", program)
```

Upload file to your flipper zero ir folder

## How it works?
A very good analysis of the protocol can be found here: https://irforum.analysir.com/viewtopic.php?t=1051&sid=7af542638f675d86e97595b78799c66e&start=10

I've modified the decoding a little bit - swapped 0 and 1 interpretation and found out, that every command includes a counter. Additionaly the last message includes the total cound of commands.

## State

Simple movements are working. Not all commands are currently implemented


## Contributing

Contributions are welcome! If you'd like to improve the project or add new features, feel free to open a pull request.

## Resources

* https://irforum.analysir.com/viewtopic.php?t=1051&sid=7af542638f675d86e97595b78799c66e&start=10
* https://botleybot.com/

## License

This project is licensed under the MIT License - see the LICENSE file for details.

Happy hacking!