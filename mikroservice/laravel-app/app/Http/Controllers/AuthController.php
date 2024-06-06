<?php

// app/Http/Controllers/AuthController.php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;

class AuthController extends Controller
{
    public function showLoginForm()
    {
        return view('auth.login');
    }

    public function login(Request $request)
    {
        $credentials = $request->only('email', 'password');

        if (Auth::attempt($credentials)) {
            // Jika berhasil login, tentukan peran pengguna
            $user = Auth::user();
        
            if ($user->role === 'admin') {
                // Jika pengguna adalah admin, arahkan ke dashboard admin
                return redirect('/dashboard/admin');
            } else {
                // Jika pengguna bukan admin, arahkan ke dashboard biasa
                return redirect('/');
            }
        } else {
            // Jika gagal login
            return redirect()->route('loginForm')->with('error', 'Email atau password salah.');
        }
        
    }

    public function showRegisterForm()
    {
        return view('auth.register');
    }

    public function register(Request $request)
    {
        $request->validate([
            'nama' => 'required|string|max:255',
            'email' => 'required|string|email|max:255|unique:users',
            'password' => 'required|string|min:8|confirmed',
        ]);

        $user = User::create([
            'nama' => $request->nama,
            'email' => $request->email,
            'password' => Hash::make($request->password),
            'role' => 'pasien', // Set default role sebagai 'pasien'
        ]);

        // Redirect ke halaman login setelah berhasil register
        return redirect()->route('loginForm')->with('success', 'Akun berhasil dibuat. Silakan login.');
    }

    public function logout()
    {
        Auth::logout();
        return redirect()->route('loginForm')->with('success', 'Anda telah berhasil logout.');
    }
}
