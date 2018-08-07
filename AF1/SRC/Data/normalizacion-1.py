from matplotlib import pyplot as plt
from PIL import Image
import time
import numpy as np

"""Normalizacion maxima"""
def maximo(im):
    tiempoIn = time.time()
    ruta = ("C:/Users/Mictam/Documents/Vision/" + im)
    im = Image.open(ruta)
    im.show()
    im3 = im
    i = 0
    while i < im3.size[0]:
        j = 0
        while j <im3.size[1]:
            maximo = max(im3.getpixel((i,j)))
            pixel = tuple([maximo, maximo, maximo])
            im3.putpixel((i,j), pixel)
            j+=1
        i+=1
    print("El Nivel Maximo De Gris Es: ", maximo)
    im3.show()  
    tiempoFin = time.time()
    print('El Proceso Tardo: ', tiempoFin - tiempoIn, 'Segundos')
    
    
"""Normalizacion minima"""    
def minimo(im):
    tiempoIn = time.time()
    ruta = ("C:/Users/Mictam/Documents/Vision/" + im)
    im = Image.open(ruta)
    im.show()
    im4 = im
    i = 0
    while i < im4.size[0]:
        j = 0
        while j < im4.size[1]:
            minimo = min(im4.getpixel((i,j)))
            pixel = tuple([minimo, minimo, minimo])
            im4.putpixel((i,j), pixel)
            j+=1
        i+=1
    print("El Nivel Minimo De Gris Es: ", minimo)
    im4.show()
    tiempoFin = time.time()
    print('El Proceso Tardo: ', tiempoFin - tiempoIn, 'Segundos')