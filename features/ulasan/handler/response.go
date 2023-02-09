package handler

import "Gurumu/features/ulasan"

type UlasanGuruResponse struct {
	ID        uint    `json:"id"`
	NamaSiswa string  `json:"nama_siswa"`
	Penilaian float32 `json:"penilaian"`
	Ulasan    string  `json:"ulasan"`
}

type AllUlasanResponse struct {
	ID        uint    `json:"id"`
	GuruId    uint    `json:"id_guru"`
	NamaGuru  string  `json:"nama_guru"`
	Penilaian float32 `json:"penilaian"`
	Ulasan    string  `json:"ulasan"`
}

func UlasanGuruToResponse(dataCore ulasan.Core) UlasanGuruResponse {
	return UlasanGuruResponse{
		ID:        dataCore.ID,
		NamaSiswa: dataCore.NamaSiswa,
		Penilaian: dataCore.Penilaian,
		Ulasan:    dataCore.Ulasan,
	}
}

func ListUlasanGuruToResponse(dataCore []ulasan.Core) []UlasanGuruResponse {
	var DataResponse []UlasanGuruResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, UlasanGuruToResponse(value))
	}
	return DataResponse
}

//All Ulasan
func AllUlasanToResponse(dataCore ulasan.Core) AllUlasanResponse {
	return AllUlasanResponse{
		ID:        dataCore.ID,
		GuruId:    dataCore.GuruId,
		NamaGuru:  dataCore.NamaGuru,
		Penilaian: dataCore.Penilaian,
		Ulasan:    dataCore.Ulasan,
	}
}

func ListAllUlasanToResponse(dataCore []ulasan.Core) []AllUlasanResponse {
	var DataResponse []AllUlasanResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, AllUlasanToResponse(value))
	}
	return DataResponse
}
