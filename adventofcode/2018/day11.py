import numpy as np
from math import floor,ceil
gserial = 1718

rack_id = np.array([[x+11 for x in range(300)] for y in range(300)]).reshape(300,300)
y = np.array([[x+1 for x in range(300)] for y in range(300)]).reshape(300,300).T
power_level = (rack_id*y+gserial)*rack_id
digit = ((power_level//100)%10)-5

def maxfilt(arr, k=3):
    x_dim,  y_dim = arr.shape
    xi, yi = floor(k/2), ceil(k/2)
    if k % 2 == 1:
        off = k - xi
    else:
        off = k//2+1
    for x in range(k//2, y_dim-k+off):
        for y in range(k//2, y_dim-k+off):
            mat = arr[x-xi:x+yi, y-xi:y+yi]
            try:
                if np.sum(mat) > largest:
                    largest = np.sum(mat)
                    serial=(y-xi+1,x-xi+1)
            except:
                largest = np.sum(mat)
                serial=(y-xi,x-xi)
    ret = serial, k, np.sum(largest)
    return ret

for x in range(1,300):
    print(maxfilt(digit,x))
