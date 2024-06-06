<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
<div class="container">
        <div class="row">
            <div class="col-md-12">
                <div class="card">
                    <div class="card-header">Edit Cuti Dokter</div>

                    <div class="card-body">
                        <form action="{{ route('cuti-dokter.update', $cutiDokter['id']) }}" method="POST">
                            @csrf
                            @method('PUT')

                            <div class="form-group">
                                <label for="tanggalMulai">Tanggal Mulai</label>
                                <input type="date" name="tanggalMulai" id="tanggalMulai" class="form-control" value="{{ $cutiDokter['tanggalMulai'] }}" required>
                            </div>

                            <div class="form-group">
                                <label for="tanggalSelesai">Tanggal Selesai</label>
                                <input type="date" name="tanggalSelesai" id="tanggalSelesai" class="form-control" value="{{ $cutiDokter['tanggalSelesai'] }}" required>
                            </div>

                            <div class="form-group">
                                <label for="keterangan">Keterangan</label>
                                <textarea name="keterangan" id="keterangan" class="form-control" required>{{ $cutiDokter['keterangan'] }}</textarea>
                            </div>

                            <button type="submit" class="btn btn-primary">Simpan</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</body>
</html>