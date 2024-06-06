<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Log;

class UlasanController extends Controller
{
    protected $apiUrl;

    public function __construct()
    {
        $this->apiUrl = 'http://192.168.154.117:8087/api/ulasan'; // Sesuaikan dengan URL servis Go Anda
    }

    public function index()
    {
        try {
            $response = Http::get($this->apiUrl . '/all');
    
            if ($response->status() == 200) {
                $ulasans = $response->json();
    
                // Periksa apakah $ulasans adalah array
                if (is_array($ulasans) && count($ulasans) > 0) {
                    return view('ulasan.index', compact('ulasans'));
                } else {
                    // Jika $ulasans bukan array atau tidak ada data ulasan
                    return view('ulasan.index', ['ulasans' => []])->with('error', 'No data available');
                }
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return view('error')->with('message', 'Server is currently down. Please try again later.');
        }
    }
    

    public function create()
    {
        return view('ulasan.create');
    }

    public function store(Request $request)
    {
        // Ambil user ID dari user yang sudah login
        $userID = Auth::id();

        // Lakukan validasi request
        $validatedData = $request->validate([
            'isiUlasan' => 'required',
        ]);

        // Format data untuk dikirim ke servis Go
        $formattedData = [
            'userID' => $userID,
            'isiUlasan' => $validatedData['isiUlasan'],
        ];

        try {
            $response = Http::post($this->apiUrl . '/create', $formattedData);

            if ($response->status() == 201) {
                return redirect()->route('ulasan.index')->with('success', 'Ulasan berhasil ditambahkan.');
            } else {
                throw new \Exception('Failed to add data to API');
            }
        } catch (\Exception $e) {
            return redirect()->route('ulasan.index')->with('error', $e->getMessage());
        }
    }

    public function edit($id)
    {
        try {
            $response = Http::get($this->apiUrl . '/' . $id);

            if ($response->status() == 200) {
                $ulasan = $response->json();
                return view('ulasan.edit', compact('ulasan'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('ulasan.index')->with('error', $e->getMessage());
        }
    }

    public function update(Request $request, $id)
    {
        // Lakukan validasi request
        $validatedData = $request->validate([
            'isiUlasan' => 'required',
        ]);

        // Format data untuk dikirim ke servis Go
        $formattedData = [
            'isiUlasan' => $validatedData['isiUlasan'],
        ];

        try {
            $response = Http::put($this->apiUrl . '/update/' . $id, $formattedData);

            if ($response->status() == 200) {
                return redirect()->route('ulasan.index')->with('success', 'Ulasan berhasil diperbarui.');
            } else {
                throw new \Exception('Failed to update data to API');
            }
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage())->withInput();
        }
    }

    public function destroy($id)
    {
        try {
            $response = Http::delete($this->apiUrl . '/delete/' . $id);

            if ($response->status() == 200) {
                return redirect()->route('ulasan.index')->with('success', 'Ulasan berhasil dihapus.');
            } else {
                throw new \Exception('Failed to delete data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('ulasan.index')->with('error', $e->getMessage());
        }
    }
}
