package ejacobson.aoc;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;

public class Main {
    public static void main(String[] args) {
        try {
            Stream<String> lines = Files.lines(Path.of("src/main/resources/input.txt"));
            List<String> input = lines.toList();
            lines.close();

            System.out.println("Part 1: ");
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}