package ejacobson.aoc;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;

public class Main {
    static final int RED = 12;
    static final int GREEN = 13;
    static final int BLUE = 14;

    public static void main(String[] args) {
        try {
            Stream<String> lines = Files.lines(Path.of("src/main/resources/input.txt"));
            List<String> input = lines.toList();
            lines.close();

            System.out.println("Part 1: " + getSumOfEligibleIds(input));
            System.out.println("Part 2: " + getPowerOfFewestCubes(input));
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    // part 1
    static int getSumOfEligibleIds(List<String> input) {
        int sum = 0;

        for (String line : input) {
            boolean eligible = true;
            int id = Integer.parseInt(line.substring(5, line.indexOf(':')));
            for (String color : line.substring(line.indexOf(':') + 2).split("; ")) {
                for (String value : color.split(", ")) {
                    int count = Integer.parseInt(value.substring(0, value.indexOf(' ')));

                    if (value.contains("red")) {
                        if (count > RED) {
                            eligible = false;
                        }
                    }
                    if (value.contains("blue")) {
                        if (count > BLUE) {
                            eligible = false;
                        }
                    }
                    if (value.contains("green")) {
                        if (count > GREEN) {
                            eligible = false;
                        }
                    }
                }
            }
            if (eligible) {
                sum += id;
            }
        }
        return sum;
    }

    // part 2
    static int getPowerOfFewestCubes(List<String> input) {
        int sum = 0;

        for (String line : input) {
            int red = 0;
            int green = 0;
            int blue = 0;
            for (String color : line.substring(line.indexOf(':') + 2).split("; ")) {
                for (String value : color.split(", ")) {
                    int count = Integer.parseInt(value.substring(0, value.indexOf(' ')));
                    if (value.contains("red")) {
                        if (count > red) {
                            red = count;
                        }
                    }
                    if (value.contains("blue")) {
                        if (count > blue) {
                            blue = count;
                        }
                    }
                    if (value.contains("green")) {
                        if (count > green) {
                            green = count;
                        }
                    }
                }
            }
            sum += red * green * blue;
        }
        return sum;
    }
}

