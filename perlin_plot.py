import matplotlib.pyplot as plt
import pandas as pd
import sys

def rgb_from_noise_value(noise_value):
	""" Translates the noise value which is from -1 to 1 into an RGB value. """
	greyscale = ((noise_value + 1.0) / 2.0) * 250
	return [greyscale, greyscale, greyscale]

def plot_perlin_noise(perlin_csv):
	""" Creates a plot representing perlin noise data from a csv. """
	noise = pd.read_csv(perlin_csv)
	noise_matrix = []

	for index, row in noise.iterrows():

		if row['y'] == 0.00:
			matrix_row = []

		rgb_value = rgb_from_noise_value(row['noise'])
		matrix_row.append(rgb_value)

		if row['y'] == 5.99:
			noise_matrix.append(matrix_row)

	plt.imshow(noise_matrix, interpolation='None')
	plt.show()

if __name__=="__main__":
	if len(sys.argv) == 2:
		plot_perlin_noise(sys.argv[1])
	else:
		print("Usage: python perlin_plot.py [perlin_csv_file]")

