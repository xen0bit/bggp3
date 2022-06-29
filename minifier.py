from ast import Bytes
from tqdm import tqdm

input = bytes([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])

# with open('crash.mp3', 'rb') as f:
#     input = f.read()

inputlen = len(input)
inputba = bytearray(input)

i = 0

composite_list = [inputba[x:x+2] for x in range(0, len(inputba),2)]

print(composite_list)