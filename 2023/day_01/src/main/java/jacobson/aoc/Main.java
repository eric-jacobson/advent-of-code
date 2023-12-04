package jacobson.aoc;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;

public class Main {
    public static void main(String[] args) {

        try {
            Stream<String> lines = Files.lines(Path.of("src/main/resources/input.txt"));
            List<String> input = lines.toList();
            lines.close();

            System.out.println("Part 1 = " + getSum(input, false));
            System.out.println("Part 2 = " + getSum(input, true));
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static int getSum(List<String> input, boolean replaceNumericWord) {
        int sum = 0;
        for (String line : input) {
            StringBuilder num = new StringBuilder();

            if (replaceNumericWord) {
                line = replace(line);
            }

            for (char c : line.toCharArray()) {
                if (Character.isDigit(c)) {
                    num.append(c);
                    break;
                }
            }

            line = new StringBuilder(line).reverse().toString();
            for (char c : line.toCharArray()) {
                if (Character.isDigit(c)) {
                    num.append(c);
                    break;
                }
            }
            if (!num.toString().isEmpty()) {
                System.out.println(num);
                sum += Integer.parseInt(num.toString());
            }
        }
        return sum;
    }

    private static String replace(String line) {
        Map<String, String> numericWords = new HashMap<>();
        numericWords.put("one", "1");
        numericWords.put("two", "2");
        numericWords.put("three", "3");
        numericWords.put("four", "4");
        numericWords.put("five", "5");
        numericWords.put("six", "6");
        numericWords.put("seven", "7");
        numericWords.put("eight", "8");
        numericWords.put("nine", "9");

        // credit to https://github.com/SAMURAii-7/AdventOfCode-2023/blob/main/Day1/Trebuchet2.java for helping me figure out
        // a bug with my replacement algorithm
        int index = Integer.MAX_VALUE;
        while (index == Integer.MAX_VALUE) {
            String first = "";
            for (Map.Entry<String, String> entry : numericWords.entrySet()) {
                int pos = line.indexOf(entry.getKey());
                if (pos != -1 && pos < index) {
                    index = line.indexOf(entry.getKey());
                    first = entry.getKey();
                }
            }
            index = -1;
            if (numericWords.containsKey(first)) {
                line = line.replace(first.substring(0, first.length() - 1), numericWords.get(first));
                index = Integer.MAX_VALUE;
            }
        }
        return line;
    }
}