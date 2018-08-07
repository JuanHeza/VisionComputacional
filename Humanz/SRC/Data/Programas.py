import cv2
import numpy as np
import Tkinter 
import tkFileDialog
import random


#Funciones
def Guardar():
	print '\n(i) - Guardar imagen resultante? \a'
	print '\n      [ S / N ]   \a'
	
	Opcion = raw_input('     ->')

	if(Opcion == 'S' or Opcion == 's'):
		try:
			print '\n(i) - Que nombre le damos a la imagen?\a'
			NombreArchivo = raw_input('      ->')
			NombreArchivo = NombreArchivo+'.png'
			cv2.imwrite(NombreArchivo,Pr_img)
			print '\n(i) - Listo!\a'
		except:
			print '\n(X) - No se pudo!\a'

	return 0

def Grises(): #Imagen se convierte a escala de grises
	#Explicacion
	print "\n\n|------------------- A Escala de grises ------------------|"
	print "|    (i) - Se convertira la imagen a escala de grises     |"
	print "|          Formula:                                       |"
	print "|            Gris = R*.299 + G*.587 + B*114               |"
	print "|           Por cada pixel                                |"
	print "|            Pixel[x,y] = Gris                            |"
	print "|           Donde:                                        |"
	print "|           R,G,B = Tonos de los colores primarios (0-255)|"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'
	for y in range(height):
		for x in range(width):
			gris = (Or_img[y,x,0]*.114) + (Or_img[y,x,1]*.587) + (Or_img[y,x,2]*.299) #se obtiene el tono de gris del pixel
			Pr_img[y,x,0] = gris; 	#azul	
			Pr_img[y,x,1] = gris; 	#verde
			Pr_img[y,x,2] = gris; 	#rojo
	print '\n(i) - Listo!\a'
	return Pr_img

def DeCabeza(): #Imagen se voltea de cabeza
    #Explicacion
	print "\n\n|------------------- Imagen de cabeza ----------------------|"
	print "|    (i) - Se mostrara la imagen de cabeza                  |"
	print "|          Se crea una copia de la imagen,se                |"
	print "|          accede a sus valores con NuevoPixel[x,y]         |"
	print "|          Formula:                                         |"
	print "|            Por cada NuevoPixel[x,y]                       |"
	print "|            NuevoPixel[x,y] = Pixel[(X-x),(Y-y)]           |"
	print "|            Donde:                                         |"
	print "|            X = Total de columnas de la matriz             |"
	print "|            Y = Total de filas de la matriz                |"
	print "|            NuevoPixel[x,y] = Pixel de la copia imagen     |"
	print "|-----------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'
	for y in range(height):
		for x in range(width):
			Pr_img[y,x] = Or_img[height-y,width-x]	#se le da el valor del pixel que se encuentra en el lado opuesto con respecto a x,y
	print '\n(i) - Listo!\a'
	return Pr_img


def Reflejo(): #Se muestra el reflejo de la imagen
	#Explicacion
	print "\n\n|------------------- Reflejo de imagen ---------------------|"
	print "|    (i) - Se mostrara el reflejo de la imagen              |"
	print "|          Se crea una copia de la imagen,se                |"
	print "|          accede a sus valores con NuevoPixel[x,y]         |"
	print "|          Formula:                                         |"
	print "|            Por cada NuevoPixel[x,y]                       |"
	print "|            NuevoPixel[x,y] = Pixel[x,(Y-y)]               |"
	print "|            Donde:                                         |"
	print "|            Y = Total de filas de la matriz                |"
	print "|            NuevoPixel[x,y] = Pixel de la copia imagen     |"
	print "|          Al final, ambas matrices/imagenes se concatenan  |"
	print "|-----------------------------------------------------------|"
	#Proceso
	print '\n(i) - Cargando...'

	for y in range(height):
		for x in range(width):
			Pr_img[y,x] = Or_img[height-y,x] #Se le da el valor del pixel que se encuentra en el lado opuesto en x
	print '\n(i) - Listo!\a'
	return np.vstack((Or_img,Pr_img))#se concatenan

def Aclarar(): #Aclara imagen
	
	#Explicacion
	print "\n\n|------------------- Aclarar Imagen ----------------------|"
	print "|    (i) - Se aclarara la imagen                          |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|           Pixel[x,y] = {R + Br, G + Br, B + B}          |"
	print "|           Donde:                                        |"
	print "|           R,G,B = Tonos de los colores primarios (0-255)|"
	print "|           Br    = El nivel de brillo                    |"
	print "|---------------------------------------------------------|"


	#Proceso

	Brillo = raw_input('\n\n(?) - Nivel de brillo a trabajar? (1-255)\n      ->')	
	Pr_img = GrisesOpenCV(); #Se pasa a grises por el metodo de open cv para facilitar lectura de pixeles
	Or_img = GrisesOpenCV(); #Se pasa a grises por el metodo de open cv para facilitar lectura de pixeles

	try:
		Brillo = int(Brillo)
		print '\n(i) - Cargando...'
		for y in range(height):
			for x in range(width):
				#Tono azul
				if(Or_img[y,x] + Brillo>255): 
					Pr_img[y,x] = 255
				else:
					Pr_img[y,x] = Or_img[y,x] + Brillo
				
		print '\n(i) - Listo!\a'
	except:
		print "(X) - Escoje un numero valido"

	return Pr_img

def InvertirColor():
	#Explicacion 
	print "\n\n|------------------- Inversion de color ------------------|"
	print "|    (i) - Se invertiran los colores  	              	  |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|           Pixel[x,y] = 255 - Intensidad del pixel       |"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'
	Or_img = GrisesOpenCV(); #Se pasa a grises por el metodo de open cv para facilitar lectura de pixeles
	Pr_img = GrisesOpenCV(); #Se pasa a grises por el metodo de open cv para facilitar lectura de pixeles

	for y in range(height):
		for x in range(width):
			Pr_img[y,x] = 255 - Or_img[y,x]
	
	print '\n(i) - Listo!\a'
	return Pr_img

def Contraste():
	#Explicacion
	print "\n\n|---------------- Elongacion de contraste ----------------|"
	print "|    (i) - Se hara una elongacion de contraste            |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|           Pixel[x,y] = m*Pixel[x,y]                     |"
	print "|          Donde: m = contraste va de 1 a 3               |"
	print "|---------------------------------------------------------|"


	m = raw_input('\n\n(?) - Nivel de contraste? (1-3)\n      ->')	
	m = float(m)

	if m > 3:
		print "\n(!) - Escojiste un valor muy grande, se asignara 3\a"
		m = 3
	elif m < 1:
		print "\n(!) - Escojiste un valor muy chico, se asignara 1\a" 

	
	#Proceso
	try:
		print '\n(i) - Cargando...'
		Pr_img = GrisesOpenCV()
		Or_img = GrisesOpenCV()	
		
		for y in range(height):
			for x in range(width):
				if(Or_img[y,x] * m>255): 
					Pr_img[y,x] = 255
				else:
					Pr_img[y,x] = int(Or_img[y,x] * m)
		print '\n(i) - Listo!\a'
	except:
		print "\n(X) - No escojiste un valor numerico.\a"

	return Pr_img

def UmbralBinario():
	#Explicacion
	print "\n\n|-------------------- Umbral Binario ---------------------|"
	print "|    (i) - Aplicando filtro de umbral binario             |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|           Si Pixel[x,y] > U :                           |"
	print "|             Pixel[x,y] = 255                            |"
	print "|           Si Pixel[x,y] <= U :                          |"
	print "|             Pixel[x,y] = 0                              |"
	print "|          Donde: U = Umbral                              |"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'

	Umbral = raw_input('\n\n(?) - Nivel de umbral a trabajar? (1-255)\n      ->')
	Umbral = int(Umbral)	
	Pr_img = GrisesOpenCV();
	Or_img = GrisesOpenCV();
	
	for y in range(height):
		for x in range(width):
			if Or_img[y,x] > Umbral:
				Pr_img[y,x] = 255
	

			if Or_img[y,x] <= Umbral:
				Pr_img[y,x] = 0

	print '\n(i) - Listo!\a'
	return Pr_img

def RuidoGaussiano():

	#Explicacion
	print "\n\n|------------------ Ruido Gaussiano ----------------------|"
	print "|                                                         |"
	print "|    (i) - Descomponiendo la imagen con RUIDO!!\a           |"
	print "|                                                         |"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Destruyendo... B^)'
	prob = .01
	thres = 1 - prob
	Pr_img = GrisesOpenCV()
	Or_img = GrisesOpenCV()

	for y in range(height):
		for x in range(width):
			rdn = random.random()

			if rdn < prob:
				Pr_img[y,x] = 0
			elif rdn > thres:
				Pr_img[y,x] = 255
			else:
				Pr_img[y,x] = Or_img[y,x]
	print '\n(i) - Listo!\a'
	return Pr_img

def FiltroMediana():
	
	print "\n\n|------------------ Filtro de la mediana------------------|"
	print "|    (i) - Aplicando filtro de umbral binario             |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|          Se obtiene su zona de vecindad                 |"
	print "|          Se ordenan las intensidades de forma ascendente|"
	print "|      Se escoje el valor que se encuentre en la mediana  |"
	print "|          Se le da el valor de la mediana al pixel       |"
	print "|                                                         |"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'
	Pr_img = GrisesOpenCV()
	Or_img = GrisesOpenCV()

	for y in range(height):
		for x in range(width):
			#Vecindad 3x3
			vecindad = [Pr_img[y,x],Pr_img[y,x-1],Pr_img[y,x+1],Pr_img[y+1,x],Pr_img[y+1,x-1],Pr_img[y+1,x+1],Pr_img[y-1,x],Pr_img[y-1,x-1],Pr_img[y-1,x+1]]
			
			#Se ordenan para saber la mediana (la mediana es ahora vecindad[4])
			vecindad.sort()
			mediana = vecindad[4]
			
			#Se obtienen las diferencias entre la mediana y el pixel actual
			difa = int(Pr_img[y,x]) - int(mediana)
			difb = int(mediana) - int(Pr_img[y,x]) 

			#Si existe una diferencia mayor a 20, se hace la correcion usando la mediana
			if difa > 20:
				Pr_img[y,x] =  mediana
			
			if difb > 20:
				Pr_img[y,x] = mediana 

	print '\n(i) - Listo!\a'
	return Pr_img

def FiltroMedia():


	print "\n\n|------------------ Filtro de la media--------------------|"
	print "|    (i) - Aplicando filtro de umbral binario             |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|          Se obtiene su zona de vecindad                 |"
	print "|          Se suman todas las intensidades y se obtiene   |"
	print "|          el promedio                                    |"
	print "|          Se le da el valor del promedio al pixel        |"
	print "|                                                         |"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'
	Pr_img = GrisesOpenCV()
	Or_img = GrisesOpenCV()

	for y in range(height):
		for x in range(width):
			#Vecindad 3x3
			vecindad = [Pr_img[y,x],Pr_img[y,x-1],Pr_img[y,x+1],Pr_img[y+1,x],Pr_img[y+1,x-1],Pr_img[y+1,x+1],Pr_img[y-1,x],Pr_img[y-1,x-1],Pr_img[y-1,x+1]]
			
			#Se obtiene el promedio
			suma = 0

			for s in range(9):
				suma = suma + vecindad[s]

			
			media = suma/9
			
			#Se obtienen las diferencias entre la media y el pixel actual
			difa = int(Pr_img[y,x]) - int(media)
			difb = int(media) - int(Pr_img[y,x]) 

			#Si existe una diferencia mayor a 20, se hace la correcion usando la mediana
			if difa > 20:
				Pr_img[y,x] =  media
			
			if difb > 20:
				Pr_img[y,x] = media

	print '\n(i) - Listo!\a'
	return Pr_img

def FiltroModa():


	print "\n\n|------------------ Filtro de la moda---------------------|"
	print "|    (i) - Aplicando filtro de umbral binario             |"
	print "|          Formula:                                       |"
	print "|          Por cada Pixel[x,y]                            |"
	print "|          Se obtiene su zona de vecindad                 |"
	print "|          Se obtiene la intensidad de moda en la vecindad|"
	print "|          Se le da el valor de la moda al pixel          |"
	print "|                                                         |"
	print "|---------------------------------------------------------|"

	#Proceso
	print '\n(i) - Cargando...'
	Pr_img = GrisesOpenCV()
	Or_img = GrisesOpenCV()

	for y in range(height):
		for x in range(width):
			#Vecindad 3x3
			vecindad = [Pr_img[y,x],Pr_img[y,x-1],Pr_img[y,x+1],Pr_img[y+1,x],Pr_img[y+1,x-1],Pr_img[y+1,x+1],Pr_img[y-1,x],Pr_img[y-1,x-1],Pr_img[y-1,x+1]]
			
			#Se obtiene la moda
			moda = 0
			moda_cant = 0
			
			for v in range(9):
				temp_cant = 0

				for w in range(9):
					if vecindad[v] == vecindad[w]:
						temp_cant = temp_cant + 1
						#termina conteo

				if temp_cant > moda_cant:
					moda = vecindad[v]
					moda_cant = temp_cant


				Pr_img[y,x] =  moda

	print '\n(i) - Listo!\a'
	return Pr_img

def FiltroOrdenRango():
	Pr_img = GrisesOpenCV()
	Or_img = GrisesOpenCV()
	
	print '\n(i) - Que rango usamos ? (0-8) \a'
	r = raw_input('     ->')
	r = int(r)

	try:		
		for y in range(height):
			for x in range(width):
				#Vecindad 3x3
				vecindad = [Pr_img[y,x],Pr_img[y,x-1],Pr_img[y,x+1],Pr_img[y+1,x],Pr_img[y+1,x-1],Pr_img[y+1,x+1],Pr_img[y-1,x],Pr_img[y-1,x-1],Pr_img[y-1,x+1]]
				
				#Se ordenan para saber la mediana (la mediana es ahora vecindad[4])
				vecindad.sort()
				
				#Se obtienen las diferencias entre la r y el pixel actual
				difa = int(Pr_img[y,x]) - int(vecindad[r])
				difb = int(vecindad[r]) - int(Pr_img[y,x]) 

				#Si existe una diferencia mayor a 20, se hace la correcion usando la mediana
				if difa > 20:
					Pr_img[y,x] =  vecindad[r]
				
				if difb > 20:
					Pr_img[y,x] = vecindad[r]
	except:
		print '\n(i) - Valor no admitido. \a'

	return Pr_img


def ConvolucionBW():
	#Proceso

	#Se relee la imagen y se carga en grises.... la diferencia con respecto a el filtro propio de grsies es que aqui solo nos encargamos de un valor y no de 3.
	Pr_img = cv2.imread(ruta,0)
	Or_img = cv2.imread(ruta,0)

	#Se crea la matriz kernel
	#Kernel = np.array([[0,-1,0],[-1,5,-1],[0,-1,0]]) #Enfocado/Sharpen
	Kernel = np.array([[0,1,0],[1,-8,1],[0,1,0]]) #Deteccion de Bordes
	
	print '\n(i) - Cargando...'
	for y in range(height):
		for x in range(width):
			PI = (Kernel[0,0]*Or_img[y-1,x-1]) + (Kernel[0,1]*Or_img[y-1,x]) + (Kernel[0,2]*Or_img[y-1,x+1])
			PI = PI + (Kernel[1,0]*Or_img[y,x-1]) + (Kernel[1,1]*Or_img[y,x]) + (Kernel[1,2]*Or_img[y,x+1])
			PI = PI + (Kernel[2,0]*Or_img[y+1,x-1]) + (Kernel[2,1]*Or_img[y+1,x]) + (Kernel[2,2]*Or_img[y+1,x+1])
			Pr_img[y,x] = PI

	print '\n(i) - Listo!\a'
	return Pr_img
			

def GrisesOpenCV():
	Pr_img = cv2.imread(ruta,0)
	return Pr_img
	
			
#Inicia programa
print "\n[---- Tarea Fundamental 1 de Vision Computacional ----] \n\n (i) - Vamos a procesar una imagen primero escoje un archivo de imagen valido "

#Seleccion de imagen
Tkinter.Tk().withdraw() 
ruta = tkFileDialog.askopenfilename(filetypes=(("Archivos JPG", "*.jpg"),
											("Archivos JPEG", "*.jpeg"),
											("Archivos PNG", "*.png"),
                                            ("Todo los archivos", "*.*")))

#Se lee la imagen
try:

	Or_img 	= cv2.imread(ruta) #Imagen original
	Pr_img	= cv2.imread(ruta) #Imagen procesada
	prop 	= Or_img.shape   #obtiene propiedades de la imagen
	height 	= prop[0] - 1 #obtiene filas
	width 	= prop[1] - 1 #obtiene cols	
	print "\n\n (i) - Imagen leida exitosamente.\n       Detalles:\n       -Ancho  -> " + str(width) + "px\n       -Altura -> " + str(height) + "px\n       -Ruta   -> " + ruta

	#Se muestra el menu

	listaFunciones = { 'A grises': Grises, 'De cabeza': DeCabeza, 'Reflejar': Reflejo, 'Aclarar': Aclarar,'Invertir': InvertirColor,'Contraste': Contraste,'Umbral Binario': UmbralBinario,'Ruido':RuidoGaussiano,'Filtro Mediana':FiltroMediana,'Filtro Media':FiltroMedia,'Filtro Moda':FiltroModa,'Filtro ORango':FiltroOrdenRango,'Convolucion':ConvolucionBW}

	print("\n\n (?) - Que deseas hacer con la imagen? Escribe uno de los siguientes comandos \n       (Respetar las mayusculas y minusculas): \n        -> A grises : Convierte la imagen a escala de grises\n        -> De cabeza : Voltear de cabeza la imagen\n        -> Reflejar : Mostrar imagen con su reflejo\n        -> Aclarar : Aclara la imagen\n        -> Invertir: Se invierte el color de la imagen\n        -> Contraste: Elongar contraste\n        -> Ruido: Se le agrega ruido/efecto sal y pimienta a la imagen\n        -> Filtro Mediana : Se corrige una imagen con el filtro de mediana\n        -> Filtro Media : El filtro del promedio\n        -> Filtro Moda : El filtro moda\n        -> Umbral Binario: Se utilizara umbral binario\n        -> Filtro ORango: Filtro de Orden de Rango\n        -> Convolucion: Convolucion")
	
	#El usuario escoje una opcion

	opcion = raw_input('       -> ')

	#Se realiza la opcion seleccionada
	try:
	    Pr_img = listaFunciones[opcion]()
	    cv2.imshow("ACTIVIDAD FUNDAMENTAL 1",Pr_img)
	    Guardar()
	except:
	    print("\n\n(X) - Esa opcion no es valida")


except:
	print "\n\n(X) - NO SELECCIONASTE UN ARCHIVO VALIDO!\a"
	print "\n\n 	 Saliendo del programa...\a"

cv2.waitKey(0)
cv2.destroyAllWindows()


