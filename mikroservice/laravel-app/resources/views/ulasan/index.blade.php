<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Daftar Ulasan</title>
    <style>
        .card-container {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
            grid-gap: 20px;
        }

        .card {
            padding: 20px;
            background-color: #f9f9f9;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            transition: transform 0.3s ease;
        }

        .card:hover {
            transform: translateY(-5px);
        }

        .card-header {
            font-weight: bold;
            margin-bottom: 10px;
        }

        .card-content {
            margin-bottom: 10px;
        }

        .card-actions {
            display: flex;
            justify-content: space-between;
        }

        .edit-link,
        .delete-form button {
            padding: 5px 10px;
            cursor: pointer;
            border: none;
            border-radius: 4px;
            color: #fff;
        }

        .edit-link {
            background-color: #007bff;
        }

        .delete-form button {
            background-color: #dc3545;
        }
    </style>
</head>

<body>
    <h1>Daftar Ulasan</h1>

    @if(session('success'))
    <div style="color: green;">{{ session('success') }}</div>
    @endif

    @if(session('error'))
    <div style="color: red;">{{ session('error') }}</div>
    @endif

    <a href="{{ route('ulasan.create') }}">Tambah Ulasan</a>

    <div class="card-container">
        @if(isset($ulasans['data']) && is_array($ulasans['data']) && count($ulasans['data']) > 0)
        @foreach($ulasans['data'] as $ulasan)
        <div class="card">
            <div class="card-header">{{ $ulasan['user_id'] }}</div>
            <div class="card-content">{{ $ulasan['isi_ulasan'] }}</div>
            <div class="card-actions">
                <form class="delete-form" action="{{ route('ulasan.destroy', ['id' => $ulasan['id']]) }}" method="post">
                    @csrf
                    @method('DELETE')
                    <button type="submit" onclick="return confirm('Are you sure?')">Delete</button>
                </form>
            </div>
        </div>
        @endforeach
        @else
        <p>No data available</p>
        @endif
    </div>
</body>

</html>
