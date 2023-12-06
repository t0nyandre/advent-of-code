import java.math.BigInteger
import java.security.MessageDigest
import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.io.path.readText

/** Will get the File from the given resource folder */
private fun pathFrom(name: String, day: String) = Path("src/$day/$name.txt")

/** Reads lines from the given input txt file. */
fun readInput(name: String, day: String) = pathFrom(name, day).readLines()

/** Reads entire file from the given input as one big String. */
fun readInputText(name: String, day: String) = pathFrom(name, day).readText()

/** Converts string to md5 hash. */
fun String.md5() =
        BigInteger(1, MessageDigest.getInstance("MD5").digest(toByteArray()))
                .toString(16)
                .padStart(32, '0')

/** The cleaner shorthand for printing output. */
fun Any?.println() = println(this)
