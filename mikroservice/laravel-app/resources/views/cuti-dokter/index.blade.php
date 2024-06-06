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
                    <div class="card-header">Daftar Cuti Dokter</div>

                    <div class="card-body">
                        <a href="{{ route('cuti-dokter.create') }}" class="btn btn-primary mb-2">Tambah Cuti Dokter</a>

                        <table class="table">
                            <thead>
                                <tr>
                                    <th>ID</th>
                                    <th>Tanggal Mulai</th>
                                    <th>Tanggal Selesai</th>
                                    <th>Keterangan</th>
                                    <th>Aksi</th>
                                </tr>
                            </thead>
                            <tbody>
                                @foreach($cutiDokters as $cutiDokter)
                                    <tr>
                                        <td>{{ $cutiDokter['id'] }}</td>
                                        <td>{{ $cutiDokter['tanggalMulai'] }}</td>
                                        <td>{{ $cutiDokter['tanggalSelesai'] }}</td>
                                        <td>{{ $cutiDokter['keterangan'] }}</td>
                                        <td>
                                            <a href="{{ route('cuti-dokter.edit', $cutiDokter['id']) }}" class="btn btn-sm btn-warning">Edit</a>
                                            <form action="{{ route('cuti-dokter.destroy', $cutiDokter['id']) }}" method="POST" class="d-inline">
                                                @csrf
                                                @method('DELETE')
                                                <button type="submit" class="btn btn-sm btn-danger" onclick="return confirm('Apakah Anda yakin ingin menghapus?')">Hapus</button>
                                            </form>
                                        </td>
                                    </tr>
                                @endforeach
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</body>
</html>