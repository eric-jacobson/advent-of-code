package ejacobson.aoc;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;

public class Main {
    static final List<int[]> adjacentPositions = new ArrayList<>(List.of(
            new int[]{-1, -1}, new int[]{-1, 0}, new int[]{-1, 1},
            new int[]{0, -1}, new int[]{0, 1},
            new int[]{1, -1}, new int[]{1, 0}, new int[]{1, 1}
    ));

    public static void main(String[] args) {
        try {
            Stream<String> lines = Files.lines(Path.of("src/main/resources/input.txt"));
            char[][] matrix = lines.map(String::toCharArray).toArray(char[][]::new);
            lines.close();

            System.out.println("Part 1: " + partNumbers(matrix));
            System.out.println("Part 2: " + gearRatios(matrix));

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static boolean isSymbol(char c) {
        return !Character.isDigit(c) && c != '.';
    }

    private static int partNumbers(char[][] matrix) {
        List<Integer> partNumbers = new ArrayList<>();
        boolean candidate = false;

        for (int r = 0; r < matrix.length; r++) {
            StringBuilder sb = new StringBuilder();
            for (int c = 0; c < matrix[r].length; c++) {
                if (Character.isDigit(matrix[r][c])) {
                    sb.append(matrix[r][c]);
                    for (int[] pos : adjacentPositions) {
                        int row = r + pos[0];
                        int col = c + pos[1];
                        if (row >= 0 && row < matrix.length && col >= 0 && col < matrix[r].length && isSymbol(matrix[row][col])) {
                            candidate = true;
                            break;
                        }
                    }
                    if (c == matrix[r].length - 1 && candidate && sb.length() > 0) {
                        partNumbers.add(Integer.parseInt(sb.toString()));
                        candidate = false;
                    }
                } else {
                    if (candidate && sb.length() > 0) {
                        partNumbers.add(Integer.parseInt(sb.toString()));
                        candidate = false;
                    }
                    sb = new StringBuilder();
                }
            }
        }
        return partNumbers.stream().mapToInt(Integer::intValue).sum();
    }

    private static int gearRatios(char[][] matrix) {
        List<Integer> partNumbers = new ArrayList<>();
        Map<String, List<Integer>> gearRatios = new HashMap<>();
        boolean candidate = false;

        for (int r = 0; r < matrix.length; r++) {
            StringBuilder sb = new StringBuilder();
            String key = "";
            for (int c = 0; c < matrix[r].length; c++) {
                if (Character.isDigit(matrix[r][c])) {
                    sb.append(matrix[r][c]);
                    for (int[] pos : adjacentPositions) {
                        int row = r + pos[0];
                        int col = c + pos[1];
                        if (row >= 0 && row < matrix.length && col >= 0 && col < matrix[r].length && matrix[row][col] == '*') {
                            key = row + "," + col;
                            gearRatios.putIfAbsent(key, new ArrayList<>());
                            candidate = true;
                            break;
                        }
                    }
                    if (c == matrix[r].length - 1 && candidate && sb.length() > 0) {
                        gearRatios.get(key).add(Integer.parseInt(sb.toString()));
                        candidate = false;
                    }
                } else {
                    if (candidate && sb.length() > 0) {
                        gearRatios.get(key).add(Integer.parseInt(sb.toString()));
                        candidate = false;
                    }
                    sb = new StringBuilder();
                }
            }
        }

        for (Map.Entry<String, List<Integer>> entry : gearRatios.entrySet()) {
            if (entry.getValue().size() == 2) {
                partNumbers.add(entry.getValue().stream().reduce(1, (a, b) -> a * b));
            }
        }
        return partNumbers.stream().mapToInt(Integer::intValue).sum();
    }
}