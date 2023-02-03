import os
import numpy as np

import matplotlib.pyplot as plt

if __name__ == '__main__':
    files = ['zad1_solution', 'zad2_solution', 'zad3_solution', 'zad4_solution']
    for file in files:
        plt.clf()
        with open(file + '.txt') as f:
            lines = f.readlines()

            time_str = lines[1]
            time_str = np.array(time_str.split(", "))
            time_str = time_str[:len(time_str) - 1]

            for i in range(3, len(lines)):
                line = lines[i]
                # if line.startswith("Trapezoid"):
                if line[0].isalpha():
                    xs_str = lines[i + 1]
                    ys_str = lines[i + 2]

                    print("converting", line)
                    xs_split = np.array(xs_str.split(", "))
                    ys_split = np.array(ys_str.split(", "))

                    xs_split = xs_split[:len(xs_split) - 1]
                    ys_split = ys_split[:len(ys_split) - 1]

                    xs = xs_split.astype(float)
                    ys = ys_split.astype(float)

                    if len(xs) > len(time_str):
                        xs = xs[:len(xs) - 1]
                    plt.subplot(1, 2, 1)
                    plt.plot(time_str, xs, label=line)

                    if len(ys) > len(time_str):
                        ys = ys[:len(ys) - 1]
                    plt.subplot(1, 2, 2)
                    plt.plot(time_str, ys, label=line)
            plt.legend()
            plt.savefig(file + ".jpg")
