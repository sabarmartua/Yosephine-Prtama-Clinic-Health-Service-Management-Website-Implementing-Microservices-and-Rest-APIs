<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Daftar Antrian</title>
</head>
<body>
    <h1>Daftar Antrian</h1>

    @if (session('success'))
        <div>{{ session('success') }}</div>
    @endif

    @if (session('error'))
        <div>{{ session('error') }}</div>
    @endif

    @if(empty($antrians))
        <div>Maaf, server sedang down.</div>
    @else
        <table border="1">
            <thead>
                <tr>
                    <th>No</th>
                    <th>Kepentingan</th>
                    <th>Tanggal</th>
                    <th>Deskripsi</th>
                    <th>Nomor Antrian</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                @foreach ($antrians as $antrian)
                    <tr>
                        <td>{{ $loop->iteration }}</td>
                        <td>{{ $antrian['kepentingan'] }}</td>
                        <td>{{ $antrian['tanggal'] }}</td>
                        <td>{{ $antrian['deskripsi'] }}</td>
                        <td>{{ $antrian['nomorAntrian'] }}</td>
                        <td>
                            <form action="{{ route('antrian.destroy', $antrian['id']) }}" method="POST">
                                @csrf
                                @method('DELETE')
                                <button type="submit">Delete</button>
                            </form>
                        </td>
                    </tr>
                @endforeach
            </tbody>
        </table>
    @endif
</body>
</html>
