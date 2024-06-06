<?php

namespace App\Http\Controllers;

use Illuminate\Http\Client\ConnectionException;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class ObatController extends Controller
{
    protected $apiUrl;

    public function __construct()
    {
        $this->apiUrl = 'http://192.168.154.117:8090/api/obat'; // Sesuaikan dengan URL servis Go Anda
    }

    public function index()
    {
        try {
            $response = Http::get($this->apiUrl . '/all');

            if ($response->status() == 200) {
                $obats = $response->json();
                return view('obat.index', compact('obats'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return view('error')->with('message', 'Server is currently down. Please try again later.');
        }
    }

    public function create()
    {
        return view('obat.create');
    }

    public function store(Request $request)
    {
        // Lakukan validasi request
        $validatedData = $request->validate([
            'nama' => 'required',
            'expiredDate' => 'required|date',
            'jumlahStok' => 'required|integer', // Pastikan jumlahStok adalah integer
            'deskripsi' => 'required',
        ]);
    
        // Ubah format tanggal ke format yang diharapkan oleh API
        $formattedData = [
            'nama' => $validatedData['nama'],
            'expiredDate' => date('Y-m-d\TH:i:s\Z', strtotime($validatedData['expiredDate'])),
            'jumlahStok' => intval($validatedData['jumlahStok']), // Konversi ke integer
            'deskripsi' => $validatedData['deskripsi']
        ];
    
        try {
            $response = Http::post($this->apiUrl . '/create', $formattedData);
    
            if ($response->status() == 200) {
                return redirect()->route('obat.index')->with('success', 'Obat berhasil ditambahkan.');
            } else {
                throw new \Exception('Failed to add data to API');
            }
        } catch (\Exception $e) {
            return redirect()->route('obat.index')->with('error', $e->getMessage());
        }
    }
    
    


    public function edit($id)
    {
        try {
            $response = Http::get($this->apiUrl . '/' . $id);

            if ($response->status() == 200) {
                $obat = $response->json();
                return view('obat.edit', compact('obat'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('obat.index')->with('error', $e->getMessage());
        }
    }

    public function update(Request $request, $id)
    {
        // Lakukan validasi request
        $validatedData = $request->validate([
            'nama' => 'required',
            'expiredDate' => 'required|date',
            'jumlahStok' => 'required|integer',
            'deskripsi' => 'required',
        ]);

        $formattedData = [
            'nama' => $validatedData['nama'],
            'expiredDate' => date('Y-m-d\TH:i:s\Z', strtotime($validatedData['expiredDate'])),
            'jumlahStok' => intval($validatedData['jumlahStok']), // Konversi ke integer
            'deskripsi' => $validatedData['deskripsi']
        ];

        try {
            $response = Http::put($this->apiUrl . '/update/' . $id, $formattedData);

            if ($response->status() == 200) {
                return redirect()->route('obat.index')->with('success', 'Obat berhasil diperbarui.');
            } else {
                throw new \Exception('Failed to update data to API');
            }
        } catch (\Exception $e) {
            return redirect()->route('obat.index')->with('error', $e->getMessage());
        }
    }

    public function destroy($id)
    {
        try {
            $response = Http::delete($this->apiUrl . '/delete/' . $id);

            if ($response->status() == 200) {
                return redirect()->route('obat.index')->with('success', 'Obat berhasil dihapus.');
            } else {
                throw new \Exception('Failed to delete data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('obat.index')->with('error', $e->getMessage());
        }
    }
}
