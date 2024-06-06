<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class CutiDokterController extends Controller
{
    protected $apiUrl;

    public function __construct()
    {
        $this->apiUrl = 'http://192.168.154.117:8085/api/cutidokter'; // Sesuaikan dengan URL servis Go Anda
    }

    public function index()
    {
        try {
            $response = Http::get($this->apiUrl . '/all');

            if ($response->status() == 200) {
                $cutiDokters = $response->json();
                return view('cuti-dokter.index', compact('cutiDokters'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return view('error')->with('message', 'Server is currently down. Please try again later.');
        }
    }

    public function create()
    {
        return view('cuti-dokter.create');
    }

    public function store(Request $request)
    {
        // Lakukan validasi request
        $validatedData = $request->validate([
            'tanggalMulai' => 'required|date',
            'tanggalSelesai' => 'required|date|after_or_equal:tanggalMulai',
            'keterangan' => 'required',
        ]);

        // Ubah format tanggal ke format yang diharapkan oleh servis Go
        $formattedData = [
            'tanggalMulai' => date('Y-m-d\TH:i:s\Z', strtotime($validatedData['tanggalMulai'])),
            'tanggalSelesai' => date('Y-m-d\TH:i:s\Z', strtotime($validatedData['tanggalSelesai'])),
            'keterangan' => $validatedData['keterangan'],
        ];

        try {
            $response = Http::post($this->apiUrl . '/create', $formattedData);

            if ($response->status() == 200) {
                return redirect()->route('cuti-dokter.index')->with('success', 'Cuti dokter berhasil ditambahkan.');
            } else {
                throw new \Exception('Failed to add data to API');
            }
        } catch (\Exception $e) {
            return redirect()->route('cuti-dokter.index')->with('error', $e->getMessage());
        }
    }

    public function edit($id)
    {
        try {
            $response = Http::get($this->apiUrl . '/' . $id);

            if ($response->status() == 200) {
                $cutiDokter = $response->json();
                return view('cuti-dokter.edit', compact('cutiDokter'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('cuti-dokter.index')->with('error', $e->getMessage());
        }
    }

    public function update(Request $request, $id)
    {
        // Lakukan validasi request
        $validatedData = $request->validate([
            'tanggalMulai' => 'required|date',
            'tanggalSelesai' => 'required|date|after_or_equal:tanggalMulai',
            'keterangan' => 'required',
        ]);

        // Ubah format tanggal ke format yang diharapkan oleh servis Go
        $formattedData = [
            'tanggalMulai' => date('Y-m-d\TH:i:s\Z', strtotime($validatedData['tanggalMulai'])),
            'tanggalSelesai' => date('Y-m-d\TH:i:s\Z', strtotime($validatedData['tanggalSelesai'])),
            'keterangan' => $validatedData['keterangan'],
        ];

        try {
            $response = Http::put($this->apiUrl . '/update/' . $id, $formattedData);

            if ($response->status() == 200) {
                return redirect()->route('cuti-dokter.index')->with('success', 'Cuti dokter berhasil diperbarui.');
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
                return redirect()->route('cuti-dokter.index')->with('success', 'Cuti dokter berhasil dihapus.');
            } else {
                throw new \Exception('Failed to delete data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('cuti-dokter.index')->with('error', $e->getMessage());
        }
    }
}
