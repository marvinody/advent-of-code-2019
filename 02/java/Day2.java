import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

public class Day2 {
  // boilerplate to read stdin and pass to driver
  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);

    ArrayList<String> lines = new ArrayList<>();
    while (scanner.hasNext()) {
      String line = scanner.nextLine();
      lines.add(line);
    }
    scanner.close();
    driver(lines);
  }

  public static final int wantedOutput = 19690720;
  // processes the input and drives the main function of interest
  public static void driver(List<String> lines) {
    String[] rawInsns = lines.get(0).split(",");
    int[] insns = Arrays.asList(rawInsns)
      .stream()
      .mapToInt(Integer::parseInt)
      .toArray();

    Emulator emulator = new Emulator(insns);
    emulator.run();

    for(int noun = 0; noun < insns.length; noun++) {
      insns[1] = noun;
      for (int verb = 0; verb < insns.length; verb++) {
        insns[2] = verb;
        // make and RUN. nice bug because I wasn't running it
        Emulator bruteForceEmu = new Emulator(insns);
        bruteForceEmu.run();
        int[] state = bruteForceEmu.getState();
        if(state[0] == wantedOutput) {
          System.out.printf("p2: noun: %d, verb: %d\n", noun, verb);
        }

      }
    }
    System.out.println("Finished iterating for amount of insns");

    int[] state = emulator.getState();
    System.out.printf("Output: %d\n", state[0]);
  }

}

// will step through the insns until it hits a HALT insn
class Emulator {
  private int[] insns;
  private int PC;
  public Emulator(int[] insns) {
    this.insns = insns.clone();
  }

  // not sure where to put these but I think this is ok?
  // I don't wanna make enums because the values actually matter
  // but it seems cluttered here
  private static final int ADD = 1;
  private static final int MULTIPLY = 2;
  private static final int HALT = 99;

  public void run() {
    while (true) { // could set a max insn count or see when PC breaks OOB
      int insn = this.insns[this.PC];
      // this is the only way to get out of our while true
      if (insn == Emulator.HALT) {
        return;
      }

      int leftAddr = this.insns[this.PC + 1];
      int rightAddr = this.insns[this.PC + 2];
      int outAddr = this.insns[this.PC + 3];

      int result;
      switch (insn){
        case ADD:
          result = this.insns[leftAddr] + this.insns[rightAddr];
          insns[outAddr] = result;
          break;
        case MULTIPLY:
          result = this.insns[leftAddr] * this.insns[rightAddr];
          insns[outAddr] = result;
          break;
      }
      this.PC += 4;
    }
  }

  public int[] getState() {
    return insns;
  }

}
