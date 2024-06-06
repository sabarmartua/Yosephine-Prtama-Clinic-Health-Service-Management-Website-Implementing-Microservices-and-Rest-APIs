<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Auth;
use Illuminate\Http\Client\ConnectionException;

class AntrianController extends Controller
{
    protected $antrianApiUrl;
    protected $cutiDokterApiUrl;

    public function __construct()
    {
        $this->antrianApiUrl = 'http://192.168.154.117:8091/api/antrian'; // Sesuaikan dengan URL servis Antrian Anda
        $this->cutiDokterApiUrl = 'http://192.168.154.117:8085/api/cutidokter'; // Sesuaikan dengan URL servis Cuti Dokter Anda
    }

    public function create()
    {
        return view('antrian.create');
    }

    public function store(Request $request)
    {
        // Ambil user ID dari user yang sudah login
        $userID = Auth::id();

        // Lakukan validasi tanggal dengan data cuti dokter
        $requestedDate = $request->input('tanggal');
        $cutiDokterResponse = Http::get($this->cutiDokterApiUrl . '/all');
        $cutiDokterData = $cutiDokterResponse->json();

        $dokterLibur = false;
        foreach ($cutiDokterData as $cuti) {
            $tanggalMulai = $cuti['tanggalMulai'];
            $tanggalSelesai = $cuti['tanggalSelesai'];

            if ($requestedDate >= $tanggalMulai && $requestedDate <= $tanggalSelesai) {
                $dokterLibur = true;
                break;
            }
        }

        if ($dokterLibur) {
            return back()->with('error', 'Mohon maaf, dokter sedang cuti pada tanggal yang Anda pilih.');
        }

        // Lakukan validasi request
        $validatedData = $request->validate([
            'kepentingan' => 'required',
            'tanggal' => 'required',
            'deskripsi' => 'required',
        ]);

        // Menambahkan data pengguna yang sudah login
        $validatedData['user_id'] = $userID;
        try {
            // Coba terhubung ke server antrian
            $response = Http::post($this->antrianApiUrl . '/create', $validatedData);
    
            if ($response->successful()) {
                return redirect()->route('antrian.index')->with('success', 'Antrian berhasil ditambahkan.');
            } else {
                throw new \Exception('Failed to add data to API');
            }
        } catch (ConnectionException $e) {
            // Tangkap jika terjadi koneksi gagal dengan server antrian
            return back()->with('error', 'Mohon maaf, server antrian sedang tidak dapat diakses. Silakan coba lagi nanti.');
        } catch (\Exception $e) {
            return redirect()->route('antrian.index')->with('error', $e->getMessage());
        }
    }

    public function index()
    {
        try {
            $response = Http::get($this->antrianApiUrl . '/all');

            if ($response->status() == 200) {
                $antrians = $response->json();
                return view('antrian.index', compact('antrians'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('antrian.index')->with('error', $e->getMessage());
        }
    }

    public function destroy($id)
    {
        try {
            $response = Http::delete($this->antrianApiUrl . '/delete/' . $id);

            if ($response->status() == 200) {
                return redirect()->route('antrian.index')->with('success', 'Antrian berhasil dihapus.');
            } else {
                throw new \Exception('Failed to delete data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('antrian.index')->with('error', $e->getMessage());
        }
    }
}
