/*
Package log provides simple leveled logging.

Usage

To use this package get the logger using:

	logger = log.NewLogger(os.Stdout)

You can also get a leveled logger using:

	Logger = log.NewLogger(os.Stdout).Level(6)

(Note that we are passing Level as 6, which will log everything)

Once you have the logger you can use it to log stuff like:

	logger.Log("Hello world!")

	logger.Inform("Someone knocked!")

	logger.Alert("Intruder!")

	logger.Warn("Threat found.")

Or for formatted logging:

	logger.Logf("Hello %s!", "John")

	logger.Informf("%s knocked!", "John")

	logger.Alertf("Intruder of level %d", 9)

	logger.Warnf("Found %q threat!", "Impersonator")

*/

package log
