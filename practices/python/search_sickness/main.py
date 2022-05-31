# Modulos
import json

## Variables

# Listado de sintomas a buscar
symptoms = []
# Cantidad de sintomas a buscar
quantity_symptoms = 0
# Objeto que contiene la información sobre las posibles enfermedades.
file_sickness = open("sickness.json", "r+").read()
data_sickness = json.loads(file_sickness)
# Contador
count = 0
# Sintomas no encontrados
symptoms_not_found = []
# Resultado búsqueda
results = []

## Variables

# Se solicita la cantidad de sintomas al usuario. 
quantity_symptoms = int(input("Digite la cantidad de sintomas que presenta el paciente: "))

# Se valida que se ingrese al menos un sintoma.
while quantity_symptoms <= 0:
    print("Por favor ingrese al menos un sintoma para continuar...")
    quantity_symptoms = int(input("Digite la cantidad de sintomas que presenta el paciente: "))

# Se solicita la cantidad de sintomas ingresados.
while quantity_symptoms > len(symptoms):
    symptom = input("Ingrese el nombre del síntoma que presenta: ").lower()
    symptoms.append(symptom)

print(f"La cantidad de sintomas ingresados [{quantity_symptoms}] y son {symptoms}")

# Se recorre cada sintoma ingresado por el usuario.
for symptom_search in symptoms:
    # Se recorre cada enfermedad de la base de datos (JSON) para buscar en cada una, si el sintoma coincide.
    for sickness in data_sickness["sickness"]:
        if symptom_search in sickness["symptoms"]:
            # Se añade la enfermedad a los resultados sin repetición.
            if sickness["name"] not in results:
                results.append(sickness["name"])
            
            # Se verifica que el sintoma no haya sido añadido en el listado de sintomas no encontrados en la base de datos (JSON).
            if symptom_search in symptoms_not_found:
                symptoms_not_found.remove(symptom_search)
        else:
            if symptom_search not in symptoms_not_found:
                symptoms_not_found.append(symptom_search)

# Si hubieron sintomas que no coincideron con la base de datos (JSON), se añaden a esta.
if len(symptoms_not_found) > 0:        
    
    with open("sickness.json", "r+") as jsonFile:
        data_sickness = json.load(jsonFile)

        # Se añade al archivo los sintomas que no fueron encontrados
        for symptom in symptoms_not_found:

            data_sickness["symptoms"].append(symptom)

        jsonFile.seek(0)
        json.dump(data_sickness, jsonFile)
        jsonFile.truncate()

# Se muestra el listado de posibles enfermedades
if results:
    print(f"Se pueden presentar las siguientes enfermedades con base en los sintomas indicados: {results}")

print("Fin del programa...")