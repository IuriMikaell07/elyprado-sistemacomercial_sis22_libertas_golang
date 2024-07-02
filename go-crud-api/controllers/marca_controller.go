package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMarcas(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idmarca, nomemarca, logo, pais_origem, telefone_sac FROM marca")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var Marcas []models.Marca
	for rows.Next() {
		var marca models.Marca
		if err := rows.Scan(&marca.IDmarca, &marca.Nomemarca, &marca.Logo, &marca.Pais_origem, &marca.Telefone_sac); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Marcas = append(Marcas, marca)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Marcas)
}

func GetMarca(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var marca models.Marca
	err = db.QueryRow("SELECT idmarca, nomemarca, logo, pais_origem, telefone_sac FROM marca WHERE idmarca = ?", id).Scan(&marca.IDmarca, &marca.Nomemarca, &marca.Logo, &marca.Pais_origem, &marca.Telefone_sac)
	if err == sql.ErrNoRows {
		http.Error(w, "Marca not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marca)
}

func CreateMarca(w http.ResponseWriter, r *http.Request) {
	var marca models.Marca
	if err := json.NewDecoder(r.Body).Decode(&marca); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO marca (nomemarca, logo, pais_origem, telefone_sac) VALUES (?, ?, ?, ?)", marca.Nomemarca, marca.Logo, marca.Pais_origem, marca.Telefone_sac)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	marca.IDmarca = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marca)
}

func UpdateMarca(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var marca models.Marca
	if err := json.NewDecoder(r.Body).Decode(&marca); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE marca SET nomemarca = ?, logo = ?, pais_origem = ?, telefone_sac = ? WHERE idmarca = ?", marca.Nomemarca, marca.Logo, marca.Pais_origem, marca.Telefone_sac, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marca.IDmarca = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marca)
}

func DeleteMarca(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDmarca, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM marca WHERE idmarca = ?", IDmarca)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
